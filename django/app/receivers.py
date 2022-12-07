from .models import Player, Team, Match, Action
from django.db.models.signals import post_save, pre_save
from django.dispatch import receiver

@receiver(post_save, sender=Player)
def publish_player_created(sender, instance: Player, created: bool, **kwargs):
    if created:
        print('Player created')
        # publish('player_created', instance)

@receiver(post_save, sender=Team)
def publish_team_created(sender, instance: Team, created: bool, **kwargs):
    if created:
        print('Team created')
        # publish('player_created', instance)

@receiver(post_save, sender=Match)
def publish_match_created(sender, instance: Match, created: bool, **kwargs):
    if created:
        print('Match created')
        # publish('player_created', instance)

@receiver(pre_save, sender=Match)
def get_old_match(sender, instance: Match, **kwargs):
    try:
        instance._pre_save_instance = Match.objects.get(pk=instance.pk)
    except Match.DoesNotExist:
        instance._pre_save_instance = None
    # publish('match_updated', instance)

@receiver(post_save, sender=Match)
def publish_new_match_result(sender, instance: Match, created: bool, **kwargs):
    if not created and instance._pre_save_instance and (instance._pre_save_instance.team_a_goal != instance.team_a_goal or instance._pre_save_instance.team_b_goal != instance.team_b_goal):
        print('Match result published')
        # publish('match_result', instance)

@receiver(post_save, sender=Action)
def publish_action_added(sender, instance: Action, created: bool, **kwargs):
    if created:
        print('Action created')
        # publish('action_added', instance)

