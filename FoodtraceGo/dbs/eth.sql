DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                          `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                          `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

DROP TABLE IF EXISTS `food_items`;

CREATE TABLE `food_items` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `food_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
                              `outbound_price` double NULL DEFAULT NULL,
                              `shipping_warehouse` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
                              `outbound_date` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
                              `code` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
                              `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
                              `blockchain_address` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
                              `transaction_hash` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;