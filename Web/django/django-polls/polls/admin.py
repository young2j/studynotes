from django.contrib import admin
from .models import Choice, Question

# Register your models here.
# 定义关联
class ChoiceInline(admin.StackedInline):
# class ChoiceInline(admin.TabularInline):
    model = Choice
    extra = 1

class QuestionAdmin(admin.ModelAdmin):
    # ------------定义当点击question对象时，展示的字段--------------
    # fields = ['pub_date','question_text'] # 定义要在管理后台显示的字段和顺序
    fieldsets = [ # 将字段进行分类显示，fields与fieldsets只能存在一个
        ('description', {"fields": ['question_text']}),
        ('date_info',{"fields":['pub_date'],"classes":['collapse']}) # 折叠/隐藏显示
    ]
    # ---------点击后关联的展示模型字段---------------
    inlines = [ChoiceInline]
    # ---------question对象本身显示的字段列-----------
    list_display = ('question_text','pub_date','was_published_recently')
    # ----------添加筛选项-----------------------
    list_filter = ['pub_date']
    # ----------添加搜索项-----------------------
    search_fields = ['question_text']
    
admin.site.register(Question,QuestionAdmin)