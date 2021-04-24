CREATE TABLE accounts(
    id SERIAL PRIMARY KEY NOT NULL,
    document_number VARCHAR(11) NOT NULL,
    available_credit_limit NUMERIC(1000, 2) NOT NULL
);
CREATE TABLE operation_types(
    id SERIAL PRIMARY KEY NOT NULL,
    description VARCHAR(100) NOT NULL
);

CREATE TABLE transactions(
    id SERIAL PRIMARY KEY NOT NULL,
    account_id INTEGER NOT NULL,
    operationtype_id INTEGER NOT NULL,
    amount NUMERIC(1000,2) NOT NULL,
    event_date TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts(id),
    FOREIGN KEY (operationtype_id) REFERENCES operation_types(id)
);
ALTER TABLE transactions ALTER COLUMN event_date SET DEFAULT now();
INSERT INTO operation_types(description) VALUES('COMPRA A VISTA'), ('COMPRA PARCELADA'), ('SAQUE'), ('PAGAMENTO');
