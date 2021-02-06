import django_filters

from .models import Snippet


class SnippetFilter(django_filters.rest_framework.FilterSet):

    class Meta:
        model = Snippet
        fields = {
            "kw_category": ['in'],  # kw_category
            "operator_id": ['in'],  # 操作人
            "operate_time": ['range'],  # 操作时间

        }
