CREATE TABLE documents (
    id SERIAL PRIMARY KEY,
    doc_type VARCHAR(20) NULL,
    is_blacklisted BOOLEAN  NULL,
    doc_value   VARCHAR(20) NULL,
    created_at  bigint null,
    updated_at  bigint null
);