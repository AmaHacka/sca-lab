"""Трекер задач — учебный сервис ЛР №4 (вариант 05).

Однофайловое Django-приложение. Код намеренно простой:
предмет работы — зависимости, а не логика.
"""
import sys
import uuid

from django.conf import settings

settings.configure(
    DEBUG=False,
    SECRET_KEY="lab4-not-a-secret",
    ALLOWED_HOSTS=["*"],
    ROOT_URLCONF=__name__,
    MIDDLEWARE=[],
)

from django.core.wsgi import get_wsgi_application  # noqa: E402
from django.http import Http404, JsonResponse  # noqa: E402
from django.urls import path  # noqa: E402

ITEMS = {}


def health(request):
    return JsonResponse({"status": "ok"})


def create_item(request):
    """Создаёт задачу."""
    key = uuid.uuid4().hex[:8]
    ITEMS[key] = request.GET.get("value", "")
    return JsonResponse({"id": key})


def get_item(request, key):
    """Возвращает задачу."""
    if key not in ITEMS:
        raise Http404
    return JsonResponse({"id": key, "value": ITEMS[key]})


urlpatterns = [
    path("health", health),
    path("tasks/new", create_item),
    path("tasks/<str:key>", get_item),
]

application = get_wsgi_application()

if __name__ == "__main__":
    from django.core.management import execute_from_command_line

    execute_from_command_line([sys.argv[0], "runserver", "0.0.0.0:8080", "--noreload"])
