-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tbl_articles (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tbl_articles;
-- +goose StatementEnd
