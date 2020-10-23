CREATE TABLE disbursements (
    id integer DEFAULT nextval('disbursement_id_seq'::regclass) PRIMARY KEY,
    disbursement_code character varying NOT NULL UNIQUE,
    location jsonb,
    channel_code character varying,
    status text,
    beneficiary_id_name text,
    currency text,
    amount double precision,
    description text,
    ref_number character varying,
    created_at text,
    updated_at text
);

-- Indices -------------------------------------------------------

CREATE UNIQUE INDEX disbursement_pkey ON disbursements(id int4_ops);
CREATE UNIQUE INDEX disbursement_disbursement_code_key ON disbursements(disbursement_code text_ops);
