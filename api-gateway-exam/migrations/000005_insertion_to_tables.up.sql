-- Inserting data into the users table
INSERT INTO users (id, username, email, password, first_name, last_name, bio, website, refresh_token)
VALUES
    ('f8e1c19a-6eca-411d-8cd0-30d9716db606', 'john_doe', 'john@example.com', 'password123', 'John', 'Doe', 'I am a software engineer.', 'http://www.johndoe.com', 'refresh_token_1'),
    ('a6b4964d-69b7-46bb-a81a-354b4a22512b', 'jane_smith', 'jane@example.com', 'password456', 'Jane', 'Smith', 'I love hiking and photography.', 'http://www.janesmith.com', 'refresh_token_2'),
    ('c4fc85d1-8602-45c0-b25a-681f4eb24b80', 'alice_johnson', 'alice@example.com', 'password789', 'Alice', 'Johnson', 'Passionate about art and literature.', 'http://www.alicejohnson.com', 'refresh_token_3'),
    ('3e5c18a5-ba05-4eb9-bb0c-41bc4c1cf40c', 'sam_green', 'sam@example.com', 'passwordabc', 'Sam', 'Green', 'Tech enthusiast and gamer.', 'http://www.samgreen.com', 'refresh_token_4'),
    ('1b3f3179-bef7-4cd0-8e69-87d7e3c83a80', 'emily_brown', 'emily@example.com', 'passworddef', 'Emily', 'Brown', 'Food lover and travel enthusiast.', 'http://www.emilybrown.com', 'refresh_token_5'),
    ('7e48b83c-098f-4cf5-ae3e-9c6c789f81f2', 'michael_wang', 'michael@example.com', 'passwordghi', 'Michael', 'Wang', 'Fitness freak and sports lover.', 'http://www.michaelwang.com', 'refresh_token_6'),
    ('f0a75f4c-c799-43b2-8b0d-48c07b295e0d', 'sarah_adams', 'sarah@example.com', 'passwordjkl', 'Sarah', 'Adams', 'Animal lover and environmentalist.', 'http://www.sarahadams.com', 'refresh_token_7'),
    ('09808e33-4a4e-46a4-bd3f-d2e25f245478', 'peter_jones', 'peter@example.com', 'passwordmno', 'Peter', 'Jones', 'Bookworm and movie buff.', 'http://www.peterjones.com', 'refresh_token_8'),
    ('b2233d5d-9c97-4947-bb3a-8b1ac48f0a93', 'lisa_smith', 'lisa@example.com', 'passwordpqr', 'Lisa', 'Smith', 'Nature enthusiast and gardener.', 'http://www.lisasmith.com', 'refresh_token_9'),
    ('86e821b8-f67e-4e29-940f-39f4e5964f89', 'david_taylor', 'david@example.com', 'passwordstu', 'David', 'Taylor', 'Musician and songwriter.', 'http://www.davidtaylor.com', 'refresh_token_10'),
    ('7f31e17c-589b-4147-99e1-489870a0f92b', 'olivia_wilson', 'olivia@example.com', 'passwordvwx', 'Olivia', 'Wilson', 'Fashionista and beauty blogger.', 'http://www.oliviawilson.com', 'refresh_token_11'),
    ('2a58e7a3-cb12-4b1e-867e-601276a12497', 'ryan_baker', 'ryan@example.com', 'passwordyz', 'Ryan', 'Baker', 'Coffee lover and aspiring chef.', 'http://www.ryanbaker.com', 'refresh_token_12'),
    ('d0b69f3b-2021-4d91-8e13-c243d9eb5292', 'jennifer_clark', 'jennifer@example.com', 'password123abc', 'Jennifer', 'Clark', 'Fitness enthusiast and yoga practitioner.', 'http://www.jenniferclark.com', 'refresh_token_13'),
    ('4c07a2c0-5fb0-43d4-815e-9be56954b2aa', 'jason_miller', 'jason@example.com', 'password456def', 'Jason', 'Miller', 'Tech geek and DIY enthusiast.', 'http://www.jasonmiller.com', 'refresh_token_14'),
    ('f34d9740-5244-4955-837d-1a3d3aaf3b47', 'sophia_king', 'sophia@example.com', 'password789ghi', 'Sophia', 'King', 'Travel blogger and adventure seeker.', 'http://www.sophiaking.com', 'refresh_token_15');



-- Inserting data into the posts table
INSERT INTO posts (id, user_id, content, title, likes, dislikes, views, category)
VALUES
    ('e292ca9d-d202-4aa2-a7de-487158b02dd4', 'f8e1c19a-6eca-411d-8cd0-30d9716db606', 'This is the content of post 1.', 'Post 1', 20, 5, 100, 'Technology'),
    ('2b012de1-892b-47b0-a268-61f46eb7e0c0', 'a6b4964d-69b7-46bb-a81a-354b4a22512b', 'This is the content of post 2.', 'Post 2', 15, 3, 80, 'Travel'),
    ('57c37e24-d5f7-4d36-af51-07929a2092da', 'c4fc85d1-8602-45c0-b25a-681f4eb24b80', 'This is the content of post 3.', 'Post 3', 30, 8, 120, 'Food'),
    ('8b167a0a-33b8-48e4-8c25-b00a4e1c8c86', '3e5c18a5-ba05-4eb9-bb0c-41bc4c1cf40c', 'This is the content of post 4.', 'Post 4', 25, 7, 90, 'Science'),
    ('90d9b0f7-22f4-45c1-99b8-3d60a2f18369', '1b3f3179-bef7-4cd0-8e69-87d7e3c83a80', 'This is the content of post 5.', 'Post 5', 18, 4, 110, 'Fashion'),
    ('ea229e0f-c760-4fa4-8c53-b0cc89b87589', '7e48b83c-098f-4cf5-ae3e-9c6c789f81f2', 'This is the content of post 6.', 'Post 6', 22, 6, 95, 'Health'),
    ('cf7dd94b-7468-4a91-b0b0-065d820f9e2e', 'f0a75f4c-c799-43b2-8b0d-48c07b295e0d', 'This is the content of post 7.', 'Post 7', 17, 3, 105, 'Music'),
    ('80b154fd-27e2-4ff4-9a2d-6d8e60a2b5bb', '09808e33-4a4e-46a4-bd3f-d2e25f245478', 'This is the content of post 8.', 'Post 8', 21, 5, 85, 'Art'),
    ('6b2b063b-301b-4207-8fc0-d3c8e3b10544', 'b2233d5d-9c97-4947-bb3a-8b1ac48f0a93', 'This is the content of post 9.', 'Post 9', 28, 6, 115, 'Gaming'),
    ('1e095819-8f4c-458e-b789-09f72cc665a2', '86e821b8-f67e-4e29-940f-39f4e5964f89', 'This is the content of post 10.', 'Post 10', 19, 4, 125, 'Sports'),
    ('82b9bbf1-f0fc-4ae6-b7b4-947c3e1aa179', '7f31e17c-589b-4147-99e1-489870a0f92b', 'This is the content of post 11.', 'Post 11', 23, 5, 75, 'Cooking'),
    ('7b4c6f94-982d-4c0b-bb15-342bc9f509d9', '2a58e7a3-cb12-4b1e-867e-601276a12497', 'This is the content of post 12.', 'Post 12', 16, 3, 130, 'DIY'),
    ('e2f54efb-6c0a-42fd-87a9-d69aa32fe7b7', 'd0b69f3b-2021-4d91-8e13-c243d9eb5292', 'This is the content of post 13.', 'Post 13', 20, 6, 100, 'Nature'),
    ('e4962547-6cb3-42d0-8594-301d5c9b4166', '4c07a2c0-5fb0-43d4-815e-9be56954b2aa', 'This is the content of post 14.', 'Post 14', 24, 4, 95, 'Photography'),
    ('f9fb0b8d-7b9f-4d1f-8719-874f8ec2e22d', 'f34d9740-5244-4955-837d-1a3d3aaf3b47', 'This is the content of post 15.', 'Post 15', 26, 7, 105, 'Fitness');



-- Inserting data into the comments table
INSERT INTO comments (id, post_id, user_id, content)
VALUES
    ('f8a672ec-3b49-4db8-8c44-0ac9ebcd12a2', 'e292ca9d-d202-4aa2-a7de-487158b02dd4', 'a6b4964d-69b7-46bb-a81a-354b4a22512b', 'Great post!'),
    ('720e1871-58a4-43f1-8404-3ad317623635', 'e292ca9d-d202-4aa2-a7de-487158b02dd4', 'c4fc85d1-8602-45c0-b25a-681f4eb24b80', 'Interesting read.'),
    ('8a2e22d8-208a-4221-948e-5cf7ff9e97c0', '2b012de1-892b-47b0-a268-61f46eb7e0c0', 'f8e1c19a-6eca-411d-8cd0-30d9716db606', 'Nice picture!'),
    ('67c11c31-98b3-45bb-9322-f5e7c06667d0', '8b167a0a-33b8-48e4-8c25-b00a4e1c8c86', '3e5c18a5-ba05-4eb9-bb0c-41bc4c1cf40c', 'Awesome post!'),
    ('e5989308-4cf9-475a-9305-271c35e943d3', '90d9b0f7-22f4-45c1-99b8-3d60a2f18369', '1b3f3179-bef7-4cd0-8e69-87d7e3c83a80', 'Keep it up!'),
    ('6c63b45d-8f09-41b1-960e-4387f80cc1cc', 'cf7dd94b-7468-4a91-b0b0-065d820f9e2e', 'f0a75f4c-c799-43b2-8b0d-48c07b295e0d', 'Love the music!'),
    ('0a40082a-1320-4e91-8106-e91f9d234b1a', '80b154fd-27e2-4ff4-9a2d-6d8e60a2b5bb', '09808e33-4a4e-46a4-bd3f-d2e25f245478', 'Amazing art!'),
    ('bb8247cb-df7e-4aa0-8414-8c748379a6d9', '6b2b063b-301b-4207-8fc0-d3c8e3b10544', 'b2233d5d-9c97-4947-bb3a-8b1ac48f0a93', 'Love gaming too!'),
    ('c9ef0cb5-30d9-4aa2-ae15-1a269067f303', '1e095819-8f4c-458e-b789-09f72cc665a2', '86e821b8-f67e-4e29-940f-39f4e5964f89', 'Sports are life!'),
    ('db7e9472-5d85-4a80-8112-3c30a8ecf9d7', '82b9bbf1-f0fc-4ae6-b7b4-947c3e1aa179', '7f31e17c-589b-4147-99e1-489870a0f92b', 'Yum! Share the recipe?'),
    ('effb7423-9f11-4c26-a0fc-cd83b7d11f57', '7b4c6f94-982d-4c0b-bb15-342bc9f509d9', '2a58e7a3-cb12-4b1e-867e-601276a12497', 'Great job!'),
    ('d2af71d0-2c24-49a3-9965-cbe81b24118e', 'e2f54efb-6c0a-42fd-87a9-d69aa32fe7b7', 'd0b69f3b-2021-4d91-8e13-c243d9eb5292', 'Love nature walks!'),
    ('06326f78-af06-4349-8237-8449f9b1a8d7', 'e4962547-6cb3-42d0-8594-301d5c9b4166', '4c07a2c0-5fb0-43d4-815e-9be56954b2aa', 'Stunning photography!'),
    ('d9c79955-0f82-47b8-9d7a-af6668f40e1d', 'f9fb0b8d-7b9f-4d1f-8719-874f8ec2e22d', 'f34d9740-5244-4955-837d-1a3d3aaf3b47', 'Keep up the good work!');

