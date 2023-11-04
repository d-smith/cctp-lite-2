create table if not exists attestations (
    id integer primary key autoincrement,
    spent boolean not null default false,
    nonce integer not null,
    sender text not null,
    receiver text not null,
    source_domain integer not null,
    dest_domain integer not null,
    amount integer not null,
    message text not null,
    txnid text not null
);