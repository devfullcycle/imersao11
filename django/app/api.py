from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.decorators import api_view
from .use_cases import UpdatePlayersInMyTeam


#PUT
@api_view(['PUT'])
def update_players_in_my_team(request: Request, my_team_uuid):
    UpdatePlayersInMyTeam().execute(my_team_uuid, request.data.get('players_uuid'))
    return Response(None, 204)