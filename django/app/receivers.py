from .models import Player, Team, Match, Action, MyTeam
from django.db.models.signals import post_save, pre_save
from django.dispatch import receiver
from .producer import safe_publish_message
import json


@receiver(post_save, sender=Player)
def publish_player_created(sender, instance: Player, created: bool, **kwargs):
    if created:
        print('Player created')
        safe_publish_message(
            'newPlayer',
            json.dumps({
                'id': str(instance.uuid),
                'name': instance.name,
                'initial_price': instance.initial_price,
            })
        )


@receiver(post_save, sender=Team)
def publish_team_created(sender, instance: Team, created: bool, **kwargs):
    if created:
        print('Team created')


@receiver(post_save, sender=MyTeam)
def publish_my_players_saved(sender, instance: MyTeam, created: bool, **kwargs):
    print('My players saved')
    my_team = Team.objects.get(pk=instance.id)
    safe_publish_message(
        'chooseTeam',
        json.dumps({
            # 'my_team_id': instance.id,
            'my_team_id': str(my_team.uuid),
            'players': [str(player.uuid) for player in my_team.players.all()]
        })
    )
    # publish('player_created', instance)


@receiver(post_save, sender=Match)
def publish_match_created(sender, instance: Match, created: bool, **kwargs):
    if created:
        print('Match created')
        safe_publish_message(
            'newMatch',
            json.dumps({
                'id': str(instance.uuid),
                'match_date': instance.match_date.isoformat(),
                'team_a_id': str(instance.team_a.uuid),
                'team_a_name': instance.team_a.name,
                'team_b_id': str(instance.team_b.uuid),
                'team_b_name': instance.team_b.name,
            })
        )


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
        safe_publish_message(
            'matchUpdateResult',
            json.dumps({
                'match_id': str(instance.uuid),
                'result': f"{instance.team_a_goal}-{instance.team_b_goal}",
            })
        )


@receiver(post_save, sender=Action)
def publish_action_added(sender, instance: Action, created: bool, **kwargs):
    if created:
        print('Action created')
        safe_publish_message(
            'newAction', 
            json.dumps({
                'match_id': str(instance.match.uuid),
                'team_id': str(instance.team.uuid),
                'player_id': str(instance.player.uuid),
                'minutes': instance.minutes,
                'action': instance.action
            })
        )
