DO $$
BEGIN
    CREATE TYPE GROUPS_ROLES AS ENUM ('ADMIN', 'MEMBER');
EXCEPTION
    WHEN duplicate_object THEN null;
END$$;

DO $$
BEGIN
    CREATE TYPE THEMES AS ENUM ('DARK', 'LIGHT');
EXCEPTION
    WHEN duplicate_object THEN null;
END$$;

CREATE TABLE IF NOT EXISTS links (
    id SERIAL,

    link VARCHAR NOT NULL UNIQUE,
    created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_on TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    link_id INTEGER NOT NULL UNIQUE,

    password VARCHAR NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_on TIMESTAMP,
    last_login TIMESTAMP,
    theme THEME NOT NULL DEFAULT 'LIGHT',

    PRIMARY KEY (id),
    FOREIGN KEY (link_id) REFERENCES links(id)
);

CREATE TABLE IF NOT EXISTS users_contacts (
    id SERIAL,
    user_id INTEGER NOT NULL,
    contact_id INTEGER NOT NULL,

    created_on TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (contact_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS groups (
    id SERIAL,
    link_id INTEGER NOT NULL UNIQUE,

    name VARCHAR NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_on TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    FOREIGN KEY (link_id) REFERENCES links(id)
);

CREATE TABLE IF NOT EXISTS users_groups (
    id SERIAL,
    user_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,

    role GROUPS_ROLES NOT NULL DEFAULT 'MEMBER',
    created_on TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

-- Main user
INSERT INTO
    links(link)
VALUES
    ('83c0dc7797cf77e2fe8d47a6a0e8bd07e0eed1fd7ad988a12a5ccf5e284fa336')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    users(link_id, password)
VALUES
    (1, '24061502')
ON CONFLICT DO
    NOTHING;
