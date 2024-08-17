/* ---- STORAGE_GROUP ---- */

CREATE TABLE storage_group
(
    id                INTEGER   NOT NULL PRIMARY KEY,
    group_name        TEXT      NOT NULL,
    owner             TEXT      NOT NULL REFERENCES account (address),
    source_type       TEXT      NOT NULL,
    extra             TEXT      NOT NULL,
    height            BIGINT    NOT NULL,
    tags              TEXT
);

/* ---- BUCKET ---- */

CREATE TABLE bucket
(
    id                INTEGER      NOT NULL PRIMARY KEY,
    bucket_name       TEXT         NOT NULL,
    owner             TEXT         NOT NULL REFERENCES account (address),
    visibility        TEXT         NOT NULL,
    source_type       TEXT         NOT NULL,
    create_at         TIMESTAMP    NOT NULL,
    payment_address   TEXT         NOT NULL REFERENCES account (address),
    bucket_status     TEXT         NOT NULL,
    charged_read_quota             BIGINT  NOT NULL,
    global_virtual_group_family_id INT     NOT NULL,
    height                         BIGINT  NOT NULL,
    sp_as_delegated_agent_disabled BOOLEAN,
    tags                           TEXT
);

/* ---- OBJECT ---- */

CREATE TABLE object
(
    id                      INTEGER   NOT NULL PRIMARY KEY,
    object_name             TEXT      NOT NULL,
    bucket_name             TEXT      NOT NULL,
    owner                   TEXT      NOT NULL REFERENCES account (address),
    creator                 TEXT      NOT NULL REFERENCES account (address),
    payload_size            BIGINT    NOT NULL,
    visibility              TEXT      NOT NULL,
    content_type            TEXT      NOT NULL,
    object_status           TEXT      NOT NULL,
    redundancy_type         TEXT      NOT NULL,
    source_type             TEXT      NOT NULL,
    checksums               TEXT      NOT NULL,
    create_at               TIMESTAMP NOT NULL,
    local_virtual_group_id  INT       NOT NULL,
    height                  BIGINT    NOT NULL,
    tags                    TEXT,
    is_updating             BOOLEAN,
    updated_at              TIMESTAMP,
    updated_by              TEXT,
    version                 BIGINT
);