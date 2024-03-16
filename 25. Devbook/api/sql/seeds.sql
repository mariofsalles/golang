INSERT INTO users (username, nick, email, userpass) VALUES
('Alice Moreira', 'aliceM', 'alice.moreira@email.com', '$2a$10$TK2/jMiiegtlelOckVo6beS1ccyRHQoGoUbSZ3Jvqw3fwG61.K/I2'),
('Bruno Silva', 'brunoS', 'bruno.silva@email.com', '$2a$10$TK2/jMiiegtlelOckVo6beS1ccyRHQoGoUbSZ3Jvqw3fwG61.K/I2'),
('Carla Dias', 'carlaD', 'carla.dias@email.com', '$2a$10$TK2/jMiiegtlelOckVo6beS1ccyRHQoGoUbSZ3Jvqw3fwG61.K/I2');

INSERT INTO followers (user_id, follower_id) VALUES (1, 2),(1, 3),(3, 1);
INSERT INTO followers (user_id, follower_id) VALUES (1, 2),(1, 3),(3, 1) ON CONFLICT (user_id, follower_id) DO NOTHING;

INSERT INTO posts (title, content, author_id) VALUES
('Post from user 1', 'Content from user 1', 1),
('Post from user 2', 'Content from user 2', 2),
('Post from user 3', 'Content from user 3', 3);