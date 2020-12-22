"""demosite URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/3.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include
# from rest_framework.documentation import include_docs_urls
from rest_framework_swagger.renderers import OpenAPIRenderer, SwaggerUIRenderer
# from rest_framework.schemas import get_schema_view
from rest_framework_swagger.views import get_swagger_view
# schema_view = get_schema_view(title='api docs', renderer_classes=[
#                               OpenAPIRenderer, SwaggerUIRenderer])
schema_view = get_swagger_view(title='api 文档')

urlpatterns = [
    path('admin/', admin.site.urls),
    # 当包括其它 urlpatterns 时应该总是使用 include()
    path('polls/', include('polls.urls')),
    path('docs/', schema_view),
    # path('docs/', include_docs_urls(title='api 文檔'))
]
