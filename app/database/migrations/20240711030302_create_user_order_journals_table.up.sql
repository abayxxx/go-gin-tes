CREATE TABLE IF NOT EXISTS user_order_journals(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    product_id BIGINT UNSIGNED NOT NULL,
    price DECIMAL(10,2) DEFAULT 0.00,
    quantity INT DEFAULT 0,
    last_balance_user DECIMAL(10,2) DEFAULT 0.00,
    status INT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (product_id) references products (id) on delete cascade
) collate = utf8mb4_unicode_ci;