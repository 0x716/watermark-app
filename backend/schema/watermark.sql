Create Table watermark (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    width BIGINT NOT NULL,
    height BIGINT NOT NULL,
    opacity REAL NOT NULL,
    create_at TIMESTAMP DEFAULT NOW() NOT NULL,
    update_at TIMESTAMP DEFAULT NOW() NOT NULL
);