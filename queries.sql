--  create loans table query
CREATE TABLE loan
(
    id           VARCHAR(14) PRIMARY KEY,
    user_id      VARCHAR(14) NOT NULL,
    amount       BIGINT,
    term         INT,
    status       VARCHAR(15) NOT NULL,
    approved_by  VARCHAR(50),
    disbursed_at TIMESTAMP,
    created_at   TIMESTAMP   NOT NULL,
    updated_at   TIMESTAMP   NOT NULL,
    deleted_at   TIMESTAMP
);

-- create repayments table query
CREATE TABLE repayment
(
    id          VARCHAR(14) PRIMARY KEY,
    loan_id     VARCHAR(14) NOT NULL,
    amount      BIGINT,
    paid_amount BIGINT,
    status      VARCHAR(15) NOT NULL,
    due_date    DATE,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL,
    deleted_at  TIMESTAMP
);
