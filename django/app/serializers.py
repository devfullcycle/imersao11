from rest_framework import serializers
from .models import Player

class UpdateMyPlayersSerializer(serializers.Serializer):
    players_id = serializers.PrimaryKeyRelatedField(
        many=True, queryset=Player.objects.all()
    )