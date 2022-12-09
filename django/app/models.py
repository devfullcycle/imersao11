from django.db import models
import uuid

class UUIDModel(models.Model):
    uuid = models.UUIDField(default=uuid.uuid4, editable=False, unique=True)

    class Meta:
        abstract = True

class Player(UUIDModel):
    name = models.CharField(max_length=50)
    initial_price = models.FloatField()

    def __str__(self):
        return self.name


class Team(UUIDModel):
    name = models.CharField(max_length=50)

    def __str__(self):
        return self.name

# este model é apenas decorativo
class MyTeam(UUIDModel):
    players = models.ManyToManyField(Player)

    def __str__(self):
        return [player.name for player in self.players.all()].__str__()


class Match(UUIDModel):
    match_date = models.DateTimeField()
    team_a = models.ForeignKey(
        Team, on_delete=models.PROTECT, related_name='team_a_matches')
    team_a_goal = models.IntegerField(default=0)
    team_b = models.ForeignKey(
        Team, on_delete=models.PROTECT, related_name='team_b_matches')
    team_b_goal = models.IntegerField(default=0)

    def __str__(self):
        return f"{self.team_a.name} x {self.team_b.name}"


class Action(UUIDModel):

    player = models.ForeignKey(Player, on_delete=models.PROTECT)
    team = models.ForeignKey(Team, on_delete=models.PROTECT)
    minutes = models.IntegerField()
    match = models.ForeignKey(
        Match, on_delete=models.PROTECT, related_name='actions')

    class Actions(models.TextChoices):
        GOAL = 'goal', 'Goal'
        ASSIST = 'assist', 'Assist'
        YELLOW_CARD = 'yellow card', 'Yellow Card'
        RED_CARD = 'red card', 'Red Card'

    action = models.CharField(max_length=50, choices=Actions.choices)

    def __str__(self):
        return f"{self.minutes}' - {self.player.name} - {self.action}"
