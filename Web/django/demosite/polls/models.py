from pygments.styles import get_all_styles
from pygments.lexers import get_all_lexers
import datetime
from django.db import models
from django.utils import timezone

# 字段类型及参数选项参考
# https://docs.djangoproject.com/zh-hans/3.1/ref/models/fields/


# Create your models here.
class Question(models.Model):
    question_text = models.CharField(verbose_name='问题', max_length=200)
    pub_date = models.DateTimeField(verbose_name='发布时间')

    def was_published_recently(self):
        # return self.pub_date>=timezone.now()-datetime.timedelta(days=1) #bug
        now = timezone.now()
        return now - datetime.timedelta(days=1) <= self.pub_date <= now
    # 给方法添加后台排序依据字段
    was_published_recently.admin_order_field = 'pub_date'
    was_published_recently.boolean = True
    was_published_recently.short_description = '是否最近发布'

    def __str__(self):
        return self.question_text

    class Meta:
        ordering = ['pub_date']
        # db_table= 'polls_question' # 默认表名
        # verbose_name = '这是表的备注'
        verbose_name_plural = '问题'
        # abstract = True # 抽象基类

    # 重写save方法
    def save(self, *args, **kwargs):
        # do_something()
        super().save(*args, **kwargs)  # Call the "real" save() method.
        # do_something_else()


class Choice(models.Model):
    # 默认情况下， Django 会给每一个模型添加下面的字段，其他字段指定primary_key=True会覆盖此行为
    # id = models.AutoField(primary_key=True)

    # ForeignKey指定多对一的关系，默认会添加_id后缀，to_field可指定外键字段，默认为primary_key
    question = models.ForeignKey(Question, on_delete=models.CASCADE)
    choice_text = models.CharField(verbose_name='选项', max_length=200)
    votes = models.IntegerField(verbose_name='票数', default=0)
    # 二元组选项
    SHIRT_SIZES = (
        ('S', 'Small'),
        ('M', 'Medium'),
        ('L', 'Large'),
    )
    # S、M、L存放在数据库中，要获得第二项表示值 通过instance.get_fieldname_display(),即instance.get_shirt_size_display()
    shirt_size = models.CharField(
        max_length=1, choices=SHIRT_SIZES, default='Medium')

    # 枚举类型
    MedalType = models.TextChoices('MedalType', 'GOLD SILVER BRONZE')
    medal = models.CharField(
        blank=True, choices=MedalType.choices, max_length=10)

    # json类型
    # options = models.JSONField(null=True,help_text='测试json字段') #sqlite不支持
    def __str__(self):
        return self.choice_text


# 代理模型，与父类模型操作同一张数据表，相当于只是给原来的模型附加特定功能
# 一个代理模型必须继承自一个非抽象模型类
class QuestionProxy(Question):
    class Meta:
        ordering = ["pub_date"]
        proxy = True


LEXERS = [item for item in get_all_lexers() if item[1]]
LANGUAGE_CHOICES = sorted([(item[1][0], item[0]) for item in LEXERS])
STYLE_CHOICES = sorted([(item, item) for item in get_all_styles()])


class Snippet(models.Model):
    created = models.DateTimeField(auto_now_add=True)
    title = models.CharField(max_length=100, blank=True, default='')
    code = models.TextField()
    linenos = models.BooleanField(default=False)
    language = models.CharField(
        choices=LANGUAGE_CHOICES, default='python', max_length=100)
    style = models.CharField(choices=STYLE_CHOICES,
                             default='friendly', max_length=100)

    class Meta:
        ordering = ['created']
