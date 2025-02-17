CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL, 
    last_name VARCHAR(100) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    hashed_password VARCHAR(255) NOT NULL
); 

CREATE TABLE IF NOT EXISTS user_friends (
    id_user INTEGER NOT NULL REFERENCES users,
    id_friend INTEGER NOT NULL REFERENCES users,
    PRIMARY KEY (id_user, id_friend)
);

CREATE TABLE IF NOT EXISTS chat (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chat_participants (
    id_chat INT NOT NULL,
    id_user INT NOT NULL,
    PRIMARY KEY (id_chat, id_user),
    FOREIGN KEY (id_chat) REFERENCES chat(id),
    FOREIGN KEY (id_user) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS chat_messages (
    id SERIAL PRIMARY KEY,
    id_chat INTEGER NOT NULL REFERENCES chats,
    sender VARCHAR(50) NOT NULL,
    receiver VARCHAR(50) NOT NULL,
    date_sent TIMESTAMP NOT NULL,
    content VARCHAR(255) NOT NULL
);