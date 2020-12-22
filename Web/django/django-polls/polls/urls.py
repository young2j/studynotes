from django.urls import path
from . import views

app_name = 'polls' # 路由命名空间
urlpatterns = [
    # path('',views.index,name='index'), # name 主要运用于模板中的反向解析
    # path('<int:question_id>/detail',views.detail,name='detail'),
    # path('<int:question_id>/results',views.results,name='results'),
    path('<int:question_id>/vote',views.vote,name='vote'),

    path('',views.IndexView.as_view(),name='index'),
    path('<int:pk>/detail',views.DetailView.as_view(),name='detail'),
    path('<int:pk>/results',views.ResultsView.as_view(),name='results'),

]