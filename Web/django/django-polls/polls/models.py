import datetime
from django.db import models
from django.utils import timezone

# Create your models here.
class Question(models.Model):
    question_text = models.CharField(verbose_name='问题',max_length =200)
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



class Choice(models.Model):
    question = models.ForeignKey(Question,on_delete=models.CASCADE)
    choice_text = models.CharField(verbose_name='选项', max_length=200)
    votes = models.IntegerField(verbose_name='票数' ,default=0)

    def __str__(self):
        return self.choice_text
    
