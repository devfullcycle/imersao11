from django.contrib import admin
from .models import Player, Team, MyTeam, Match, Action

class ActionInline(admin.TabularInline):
    model = Action


class MatchAdmin(admin.ModelAdmin):
    inlines = [ActionInline]


admin.site.register(Player)
admin.site.register(Team)
admin.site.register(MyTeam)
admin.site.register(Match, MatchAdmin)
admin.site.register(Action)