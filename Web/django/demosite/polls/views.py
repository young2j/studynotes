from rest_framework.parsers import JSONParser
from rest_framework import viewsets
from rest_framework import generics
from rest_framework.decorators import action, api_view

from django.views.decorators.csrf import csrf_exempt
from django.http import HttpResponse, JsonResponse
from datetime import datetime
from django.http import Http404, HttpResponse, HttpResponseRedirect
from django.shortcuts import render, get_object_or_404
from django.urls import reverse
from django.views import generic
from django.utils import timezone
from rest_framework.response import Response

from .models import Choice, Question, Snippet
from .serializers import SnippetSerializer

# Create your views here.
# 函数视图


def index(request):
    latest_question_list = Question.objects.order_by('-pub_date')[:5]
    context = {
        'latest_question_list': latest_question_list
    }
    return render(request, 'polls/index.html', context)


def detail(request, question_id):
    # try:
    #     question = Question.objects.get(pk=question_id)
    # except Question.DoesNotExist:
    #     raise Http404('Question does not exist.')
    question = get_object_or_404(Question, pk=question_id)
    return render(request, 'polls/detail.html', {'question': question})


def results(request, question_id):
    question = get_object_or_404(Question, pk=question_id)
    return render(request, 'polls/results.html', {'question': question})


def vote(request, question_id):
    question = get_object_or_404(Question, pk=question_id)
    try:
        selected_choice = question.choice_set.get(pk=request.POST['choice'])
    except (KeyError, Choice.DoesNotExist):
        return render(request, 'polls/detail.html', {
            'question': question,
            'error_message': "you didn't selected a choice"
        })
    else:
        selected_choice.votes += 1
        selected_choice.save()
        return HttpResponseRedirect(reverse('polls:results', args=(question.id,)))


# 类视图, generic为通用视图
class IndexView(generic.ListView):
    template_name = 'polls/index.html'
    # 改变默认提供的context属性名
    context_object_name = 'latest_question_list'

    def get_queryset(self):
        """Return the last five published questions."""
        return Question.objects.filter(pub_date__lte=timezone.now()).order_by('-pub_date')[:5]


class DetailView(generic.DetailView):
    # 指定作用的模型名，
    # DetailView会默认提供名为question的context，即context={'question':question}
    # ListVies会默认提供名为question_list的context，即context={'question_list':question_list}
    model = Question
    # 指定模板名，默认会使用<app name>/<model name>_detail.html的模板，即polls/question_detail.html
    template_name = 'polls/detail.html'

    def get_queryset(self):
        """
        Excludes any questions that aren't published yet.
        """
        return Question.objects.filter(pub_date__lte=timezone.now())


class ResultsView(generic.DetailView):
    model = Question
    template_name = 'polls/results.html'


@csrf_exempt
def snippet_list(request):
    """
    List all code snippets, or create a new snippet.
    """
    if request.method == 'GET':
        snippets = Snippet.objects.all()  # queryset
        # 传入queryset，需要指定参数many=True
        serializer = SnippetSerializer(snippets, many=True)
        return JsonResponse(serializer.data, safe=False)

    elif request.method == 'POST':
        data = JSONParser().parse(request)
        # instance可以是模型实例和queryset，data为JSON
        serializer = SnippetSerializer(data=data)
        if serializer.is_valid():
            serializer.save()
            return JsonResponse(serializer.data, status=201)
        return JsonResponse(serializer.errors, status=400)


@csrf_exempt
@api_view(['GET', 'PUT', 'DELETE'])
def snippet_detail(request, pk):
    """
    Retrieve, update or delete a code snippet.
    """
    try:
        snippet = Snippet.objects.get(pk=pk)
    except Snippet.DoesNotExist:
        return HttpResponse(status=404)

    if request.method == 'GET':
        serializer = SnippetSerializer(snippet)
        return JsonResponse(serializer.data)

    elif request.method == 'PUT':
        data = JSONParser().parse(request)
        serializer = SnippetSerializer(snippet, data=data)
        if serializer.is_valid():
            serializer.save()
            return JsonResponse(serializer.data)
        return JsonResponse(serializer.errors, status=400)

    elif request.method == 'DELETE':
        snippet.delete()
        return HttpResponse(status=204)


class SnippetViewSet(viewsets.ReadOnlyModelViewSet):
    queryset = Snippet.objects.all()
    serializer_class = SnippetSerializer
    filterset_fields = ['language']
    search_fields = ['style']  # 默认是包含字段
    ordering_fields = ['id']

    @action(methods=['GET'], detail=False)
    def all_snippet(self, request, *args, **kwargs):
        """
        docstring
        """
        return Response({'all_snippet': 111})
