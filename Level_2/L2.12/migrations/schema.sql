CREATE TABLE IF NOT EXISTS events (
                                      id VARCHAR(100) PRIMARY KEY ,
                                      user_id VARCHAR(50) NOT NULL ,
                                      title VARCHAR(100) NOT NULL ,
                                      start_time TIMESTAMPTZ NOT NULL  ,
                                      end_time TIMESTAMPTZ NOT NULL  ,


    --   constraint fk_orders_delivery foreign key (order_uid) references delivery(id),
    --   constraint fk_orders_payment foreign key (order_uid) references payment(transaction),
    --   constraint fk_orders_items foreign key (order_uid) references items(chrt_id)
);



