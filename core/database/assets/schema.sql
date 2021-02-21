


create table if not exists tsl_members (
    tsl_id bigint,
    member_id int2,
    signature bytea
);

create index if not exists ts_members__tsl_id on tsl_members (tsl_id);



create table if not exists tsls (
    id bigserial primary key,
    block_id bigint,
    tx_uuid uuid not null
);

create unique index if not exists tsls__tx_uuid on tsls (tx_uuid);
create index if not exists tsls__block_id on tsls (block_id);


create table if not exists claims_members (
     claim_id bigint,
     member_id int2,
     pub_key bytea
);

create index if not exists claims__members_claim_id on claims_members (claim_id);



create table if not exists claims (
    id bigserial primary key,
    block_id bigint,
    tx_uuid uuid not null
);

create unique index if not exists claims__tx_uuid on claims (tx_uuid);
create index if not exists claims__block_id on claims (block_id);


-- todo: potentially some additional table should be added for storing of binary log records
--       (needed for writing audit histories);
create table if not exists blocks (
    number bigint primary key,
    timestamp timestamptz default now(),
    prev_block_hash bytea,
    hash bytea,
    sig bytea,
    next_block_pub_key bytea
);

create unique index if not exists blocks__number_unique on blocks(number);