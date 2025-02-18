-- name: SaveUser :exec
INSERT INTO users (first_name, last_name, username, hashed_password)
VALUES ($1, $2, $3, $4);

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserByUsernameAndPassword :one
SELECT * FROM users
WHERE username = $1 AND hashed_password = $2 LIMIT 1;

-- name: GetAllChats :many
with ultima_mensagem as (
	select *, 
	row_number() over (partition by cm.id_chat order by cm.date_sent DESC) as row_number
	from chat_messages cm 
)
select cp.id_chat, u.first_name, u.username, um.sender, um.content, um.date_sent 
from chat_participants cp 
inner join chat_participants cp2 on cp.id_chat = cp2.id_chat
inner join users u on u.id = cp2.id_user 
inner join ultima_mensagem um on um.id_chat = cp.id_chat
where cp.id_user = $1 and cp2.id_user <> $1 and um.row_number = 1
order by um.date_sent desc;

-- name: GetChatMessages :many
select sender, receiver, date_sent, content 
from chat_messages cm 
where cm.id_chat = $1
order by date_sent desc;

-- name: GetFriends :many
select u.id as id, (u.first_name || ' ' || u.last_name) as name, u.username as username
from user_friends uf 
inner join users u on u.id = uf.id_friend
where uf.id_user = $1
order by name asc;

-- name: AddFriend :exec
insert into user_friends(id_user, id_friend) values ($1, $2);

-- name: GetChatParticipants :many
select id_user
from chat_participants cp 
where cp.id_chat = $1;