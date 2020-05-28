INSERT INTO
    links(link)
VALUES
    ('3c73302bfa052c1a2dea9b9b34a64c42418b95300cb3c3cad039b00ec2d62fa7')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    links(link)
VALUES
    ('a643c5c23f24053118c25cf79625800fad595198c61125a75f7b5f50fbec96fa')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    links(link)
VALUES
    ('002d8b7291927d2d3445f7bb2a5dbaa80839126487741fb607b89ca6b6e0f762')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    links(link)
VALUES
    ('2e3f5c379a0a7252944b836533ad2aad0d5aef3061ca641aeffef449f84f1f06')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    links(link)
VALUES
    ('7422a530825168a0e6e22b2c832f4eb86829217341400aef2d9a130966b1f876')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    users(link_id, password)
VALUES
    (2, '2406')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    users(link_id, password)
VALUES
    (3, '2406')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    users(link_id, password)
VALUES
    (4, '2406')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    users(link_id, password)
VALUES
    (5, '2406')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    users(link_id, password)
VALUES
    (6, '2406')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    contacts(user_id, contact_id)
VALUES
    (1, 2)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    contacts(user_id, contact_id)
VALUES
    (1, 3)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    contacts(user_id, contact_id)
VALUES
    (1, 4)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    contacts(user_id, contact_id)
VALUES
    (2, 1)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    contacts(user_id, contact_id)
VALUES
    (2, 5)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    contacts(user_id, contact_id)
VALUES
    (3, 1)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    contacts(user_id, contact_id)
VALUES
    (3, 4)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    contacts(user_id, contact_id)
VALUES
    (4, 1)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    contacts(user_id, contact_id)
VALUES
    (4, 3)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    contacts(user_id, contact_id)
VALUES
    (5, 1)
ON CONFLICT DO
    NOTHING;