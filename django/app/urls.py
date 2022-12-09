from django.urls import path
from .api import update_players_in_my_team

urlpatterns = [
    path('my-teams/<uuid:my_team_uuid>/players', update_players_in_my_team)
]
