<p style="text-align:center;font-size:28px;font-weight:bold;">elasticsearch</p>

# [doc](https://www.elastic.co/guide/en/elasticsearch/reference/7.3/search-aggregations-metrics.html)

# 安装

## 单节点安装

```shell
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.3.2
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.3.2
```

## 集群安装

```yaml
version: '2.2'
services:
  es01:
    image: daocloud.io/library/elasticsearch:7.7.1
    container_name: es01
    environment:
      - node.name=es01
      - discovery.seed_hosts=es02
      - cluster.initial_master_nodes=es01,es02
      - cluster.name=docker-cluster
      - bootstrap.memory_lock=true
      - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
    ulimits:
      memlock:
        soft: -1
        hard: -1
    volumes:
      - esdata01:/usr/share/elasticsearch/data
    ports:
      - 9200:9200
    networks:
      - esnet
  es02:
    image: daocloud.io/library/elasticsearch:7.7.1
    container_name: es02
    environment:
      - node.name=es02
      - discovery.seed_hosts=es01
      - cluster.initial_master_nodes=es01,es02
      - cluster.name=docker-cluster
      - bootstrap.memory_lock=true
      - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
    ulimits:
      memlock:
        soft: -1
        hard: -1
    volumes:
      - esdata02:/usr/share/elasticsearch/data
    networks:
      - esnet

volumes:
  esdata01:
    driver: local
  esdata02:
    driver: local

networks:
  esnet:
```

```shell
docker-compose up -d
# 报错
# [1]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]

# 解决
vim /etc/sysctl.conf
vm.max_map_count=262144
sudo sysctl -p
```



# 入门

es可以按照`REST`风格进行`curd`操作，api结构：

> host:port/index/type/id
>
> index: 文档索引
>
> type：文档类型，相当于sql中的表概念
>
> id： 文档的唯一id

## PUT

```shell
# kibana
PUT /customer/_doc/1
{
  "name": "John Doe"
}

# shell pretty-让输出更美观
curl -X PUT "localhost:9200/customer/_doc/1?pretty" -H 'Content-Type: application/json' -d'
{
  "name": "John Doe"
}
'
```

## GET

```shell
GET /customer/_doc/1
curl -X GET "localhost:9200/customer/_doc/1?pretty"

{
  "_index" : "customer",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 1,
  "_seq_no" : 0,
  "_primary_term" : 1,
  "found" : true,
  "_source" : {
    "name" : "John Doe"
  }
}
```

## BULK

```shell
curl -H "Content-Type: application/json" -XPOST "localhost:9200/bank/_bulk?pretty&refresh" --data-binary "@accounts.json"

curl "localhost:9200/_cat/indices?v"
```

# 搜索_search

## 升序

```shell
# 在bank索引中搜索所有记录，并升序排列。默认只返回10条记录
GET /bank/_search
{
  "query": { "match_all": {} },
  "sort": [
    { "account_number": "asc" }
  ]
}

curl -X GET "localhost:9200/bank/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": { "match_all": {} },
  "sort": [
    { "account_number": "asc" }
  ]
}
'
```

## 分页

```shell
# 分页需要在请求体中指定from和size参数
GET /bank/_search
{
  "query": { "match_all": {} },
  "sort": [
    { "account_number": "asc" }
  ],
  "from": 10,
  "size": 10
}
#--------
curl -X GET "localhost:9200/bank/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": { "match_all": {} },
  "sort": [
    { "account_number": "asc" }
  ],
  "from": 10,
  "size": 10
}
'
```

## 与或

### match、match_phrase

> match: OR
>
> match_phrase: AND

```shell
# 搜索address中包含mill或者lane的文档记录
GET /bank/_search
{
  "query": { "match": { "address": "mill lane" } }
}
#--------
curl -X GET "localhost:9200/bank/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": { "match": { "address": "mill lane" } }
}
'

# 搜索address中包含mill 以及 lane的文档记录
GET /bank/_search
{
  "query": { "match_phrase": { "address": "mill lane" } }
}
#--------
curl -X GET "localhost:9200/bank/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": { "match_phrase": { "address": "mill lane" } }
}
'
```

## bool查询

### must、must_not

```shell
GET /bank/_search
{
  "query": {
    "bool": {
      "must": [
        { "match": { "age": "40" } }
      ],
      "must_not": [
        { "match": { "state": "ID" } }
      ]
    }
  }
}
#--------
curl -X GET "localhost:9200/bank/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "bool": {
      "must": [
        { "match": { "age": "40" } }
      ],
      "must_not": [
        { "match": { "state": "ID" } }
      ]
    }
  }
}
'
```

### filter

```shell
GET /bank/_search
{
  "query": {
    "bool": {
      "must": { "match_all": {} },
      "filter": {
        "range": {
          "balance": {
            "gte": 20000,
            "lte": 30000
          }
        }
      }
    }
  }
}
#--------
curl -X GET "localhost:9200/bank/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "bool": {
      "must": { "match_all": {} },
      "filter": {
        "range": {
          "balance": {
            "gte": 20000,
            "lte": 30000
          }
        }
      }
    }
  }
}
'
```

## 聚合

### group_by

```shell
GET /bank/_search
{
  "size": 0, # 只返回聚合结果，不返回命中hits的源数据
  "aggs": {
    "group_by_state": { # 按字段state进行分组
      "terms": {
        "field": "state.keyword" # 指定state的值作为返回结果中buckets的key
      }
    }
  }
}
#--------
curl -X GET "localhost:9200/bank/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "size": 0,
  "aggs": {
    "group_by_state": {
      "terms": {
        "field": "state.keyword"
      }
    }
  }
}
'
```

### avg

```shell
GET /bank/_search
{
  "size": 0,
  "aggs": {
    "group_by_state": {
      "terms": {
        "field": "state.keyword"
      },
      "aggs": {
        "average_balance": {
          "avg": {
            "field": "balance"
          }
        }
      }
    }
  }
}
#-----
curl -X GET "localhost:9200/bank/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "size": 0,
  "aggs": {
    "group_by_state": {
      "terms": {
        "field": "state.keyword"
      },
      "aggs": {
        "average_balance": {
          "avg": {
            "field": "balance"
          }
        }
      }
    }
  }
}
'
#--------
{
  "took" : 14,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 1000,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  },
  "aggregations" : {
    "group_by_state" : {
      "doc_count_error_upper_bound" : 0,
      "sum_other_doc_count" : 743,
      "buckets" : [
        {
          "key" : "TX",
          "doc_count" : 30,
          "average_balance" : {
            "value" : 26073.3
          }
        },
        {
          "key" : "MD",
          "doc_count" : 28,
          "average_balance" : {
            "value" : 26161.535714285714
          }
        },
      ]
    }
  }
}
```

### order

```shell
GET /bank/_search
{
  "size": 0,
  "aggs": {
    "group_by_state": {
      "terms": {
        "field": "state.keyword",
        "order": { # 默认按照count进行排序，可以自行指定
          "average_balance": "desc"
        }
      },
      "aggs": { # 聚合
        "average_balance": { # 聚合名
          "avg": { # 聚合类型
            "field": "balance"
          }
        }
      }
    }
  }
}

#--------
curl -X GET "localhost:9200/bank/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "size": 0,
  "aggs": {
    "group_by_state": {
      "terms": {
        "field": "state.keyword",
        "order": {
          "average_balance": "desc"
        }
      },
      "aggs": {
        "average_balance": {
          "avg": {
            "field": "balance"
          }
        }
      }
    }
  }
}
'
```

