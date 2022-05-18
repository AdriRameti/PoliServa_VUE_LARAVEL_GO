USE PoliServa;
-- MySQL dump 10.13  Distrib 8.0.22, for Win64 (x86_64)
--
-- Host: localhost    Database: poliserva
-- ------------------------------------------------------
-- Server version	8.0.22

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `courts`
--

DROP TABLE IF EXISTS `courts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `courts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `id_sport` bigint unsigned NOT NULL,
  `sector` int NOT NULL,
  `price_h` int NOT NULL,
  `img` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `courts_id_sport_foreign` (`id_sport`),
  CONSTRAINT `courts_id_sport_foreign` FOREIGN KEY (`id_sport`) REFERENCES `sports` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `courts`
--

LOCK TABLES `courts` WRITE;
/*!40000 ALTER TABLE `courts` DISABLE KEYS */;
INSERT INTO `courts` VALUES (1,'2021-12-21 16:57:48','2021-12-21 16:57:48',7,4,6,'https://ucjcsportsclub.es/wp-content/uploads/2015/03/padel3-1024x780.jpg'),(2,'2021-12-21 16:57:48','2021-12-21 16:57:48',7,2,6,'https://ucjcsportsclub.es/wp-content/uploads/2015/03/padel3-1024x780.jpg'),(3,'2021-12-21 16:57:48','2021-12-21 16:57:48',4,2,25,'https://mundoentrenamiento.com/wp-content/uploads/2017/05/doble-penalti-en-futbol-sala-800x500.jpg'),(4,'2021-12-21 16:57:48','2021-12-21 16:57:48',3,4,20,'https://www.basquetplus.com/sites/default/files/main/articles/redes.jpg'),(5,'2021-12-21 16:57:48','2021-12-21 16:57:48',3,4,20,'https://www.basquetplus.com/sites/default/files/main/articles/redes.jpg'),(6,'2021-12-21 16:57:48','2021-12-21 16:57:48',6,1,50,'https://www.bbva.com/wp-content/uploads/2017/08/bbva-balon-futbol-2017-08-11.jpg'),(7,'2021-12-21 16:57:48','2021-12-21 16:57:48',1,2,8,'https://static2.abc.es/media/bienestar/2019/08/02/tenis-abecedario-kgNF--620x349@abc.jpg'),(8,'2021-12-21 16:57:48','2021-12-21 16:57:48',2,1,15,'https://raquetados.com/wp-content/uploads/2019/09/Frontenis.jpg'),(9,'2021-12-21 16:57:48','2021-12-21 16:57:48',6,2,50,'https://www.bbva.com/wp-content/uploads/2017/08/bbva-balon-futbol-2017-08-11.jpg'),(10,'2021-12-21 16:57:48','2021-12-21 16:57:48',3,1,20,'https://www.basquetplus.com/sites/default/files/main/articles/redes.jpg');
/*!40000 ALTER TABLE `courts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `migrations`
--

DROP TABLE IF EXISTS `migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `migrations` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `migrations`
--

LOCK TABLES `migrations` WRITE;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
INSERT INTO `migrations` VALUES (1,'2019_12_14_000001_create_personal_access_tokens_table',1),(2,'2021_11_23_200237_create_users_table',1),(3,'2021_11_23_200331_create_sports_table',1),(4,'2021_11_23_200444_create_courts_table',1),(5,'2021_11_23_200519_create_reservations_table',1);
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `personal_access_tokens`
--

DROP TABLE IF EXISTS `personal_access_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `personal_access_tokens` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tokenable_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tokenable_id` bigint unsigned NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `abilities` text COLLATE utf8mb4_unicode_ci,
  `last_used_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `personal_access_tokens_token_unique` (`token`),
  KEY `personal_access_tokens_tokenable_type_tokenable_id_index` (`tokenable_type`,`tokenable_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `personal_access_tokens`
--

LOCK TABLES `personal_access_tokens` WRITE;
/*!40000 ALTER TABLE `personal_access_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `personal_access_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reservations`
--

DROP TABLE IF EXISTS `reservations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reservations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `id_user` bigint unsigned NOT NULL,
  `id_court` bigint unsigned NOT NULL,
  `date` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `hini` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `hfin` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `total_price` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `reservations_id_user_foreign` (`id_user`),
  KEY `reservations_id_court_foreign` (`id_court`),
  CONSTRAINT `reservations_id_court_foreign` FOREIGN KEY (`id_court`) REFERENCES `courts` (`id`),
  CONSTRAINT `reservations_id_user_foreign` FOREIGN KEY (`id_user`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reservations`
--

LOCK TABLES `reservations` WRITE;
/*!40000 ALTER TABLE `reservations` DISABLE KEYS */;
INSERT INTO `reservations` VALUES (1,'2021-12-21 16:58:04','2021-12-21 16:58:04',1,7,'24/6/2033','20:04','08:20',83),(2,'2021-12-21 16:58:04','2021-12-21 16:58:04',1,6,'24/4/2048','17:39','14:20',25),(3,'2021-12-21 16:58:04','2021-12-21 16:58:04',1,7,'17/4/2027','02:33','14:35',41),(4,'2021-12-21 16:58:04','2021-12-21 16:58:04',1,9,'10/1/2041','08:37','22:06',74),(5,'2021-12-21 16:58:04','2021-12-21 16:58:04',1,3,'19/10/2043','17:09','07:23',97),(6,'2021-12-21 16:58:04','2021-12-21 16:58:04',1,10,'15/6/2023','12:37','03:57',70),(7,'2021-12-21 16:58:04','2021-12-21 16:58:04',1,6,'28/8/2047','07:04','05:23',7),(8,'2021-12-21 16:58:04','2021-12-21 16:58:04',1,1,'23/11/2043','05:30','03:53',10),(9,'2021-12-21 16:58:04','2021-12-21 16:58:04',1,10,'4/2/2023','14:40','21:05',52),(10,'2021-12-21 16:58:04','2021-12-21 16:58:04',1,1,'7/11/2021','12:05','15:35',18),(19,NULL,NULL,2,1,'7/1/2022','08:00','09:00',6),(20,NULL,NULL,2,1,'9/1/2022','08:00','09:00',6),(21,NULL,NULL,42,1,'11/1/2022','21:00','22:00',6),(22,NULL,NULL,42,7,'25/1/2022','14:00','15:00',8);
/*!40000 ALTER TABLE `reservations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sports`
--

DROP TABLE IF EXISTS `sports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sports` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `slug` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `img` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sports`
--

LOCK TABLES `sports` WRITE;
/*!40000 ALTER TABLE `sports` DISABLE KEYS */;
INSERT INTO `sports` VALUES (1,'2021-12-21 16:57:41','2021-12-21 16:57:41','tenis-5209','tenis','https://static2.abc.es/media/bienestar/2019/08/02/tenis-abecedario-kgNF--620x349@abc.jpg'),(2,'2021-12-21 16:57:41','2021-12-21 16:57:41','frontenis-2805','frontenis','https://raquetados.com/wp-content/uploads/2019/09/Frontenis.jpg'),(3,'2021-12-21 16:57:41','2021-12-21 16:57:41','basquet-2456','basquet','https://www.basquetplus.com/sites/default/files/main/articles/redes.jpg'),(4,'2021-12-21 16:57:41','2021-12-21 16:57:41','futbol sala-6340','futbol sala','https://mundoentrenamiento.com/wp-content/uploads/2017/05/doble-penalti-en-futbol-sala-800x500.jpg'),(5,'2021-12-21 16:57:41','2021-12-21 16:57:41','volei-2705','volei','https://www.fundeu.es/wp-content/uploads/2021/01/voleibol-webefesptwelve224491.jpg'),(6,'2021-12-21 16:57:41','2021-12-21 16:57:41','futbol-7310','futbol','https://www.bbva.com/wp-content/uploads/2017/08/bbva-balon-futbol-2017-08-11.jpg'),(7,'2021-12-21 16:57:41','2021-12-21 16:57:41','padel-7471','padel','http://manzasport.com/wp-content/uploads/2018/04/Pista-de-padel-Paquito-Navarro-by-Manzasport-1.jpeg');
/*!40000 ALTER TABLE `sports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `surnames` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mail` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `pass` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `img` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `google2fa_secret` text COLLATE utf8mb4_unicode_ci,
  `role` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_blocked` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2021-12-21 16:57:56','2021-12-21 16:57:56','Kathryn Orn','Adrian Morar','20fe1930-805b-1516-181a-22f8350ceb5a','kkreiger@example.org','$2y$10$iSCC6SmtUNYUubUijDkXJOTC.WM4UgMUXw3HfIQtxirZUIn1RAGhq','https://upload.wikimedia.org/wikipedia/commons/thumb/1/12/User_icon_2.svg/640px-User_icon_2.svg.png',NULL,'client',0),(2,'2021-12-21 16:57:56','2021-12-21 16:57:56','Kaitlyn Mohr','Chad Fay','276406c2-d554-5ab6-38b1-1c0a72820b42','lincoln15@example.net','$2y$10$x3UdhVh8dn8L4cevFCsjn..Nv7dtoDpr3NOZt7PPWKM3BuD.34cwe','https://upload.wikimedia.org/wikipedia/commons/thumb/1/12/User_icon_2.svg/640px-User_icon_2.svg.png',NULL,'client',0),(3,'2021-12-21 16:57:56','2021-12-21 16:57:56','Mr. Nickolas Schowalter','Claudie Metz','57359b37-558b-4e13-49ce-bca9baf71443','dagmar41@example.net','$2y$10$LAvVDOScNVhnpNrQ7F62ZeAEh9jzxFrXerd/0SS4gyWXleWlkuBI.','https://upload.wikimedia.org/wikipedia/commons/thumb/1/12/User_icon_2.svg/640px-User_icon_2.svg.png',NULL,'client',0),(4,'2021-12-21 16:57:56','2021-12-21 16:57:56','Jackson Wintheiser','Wilford Durgan','451bba0f-620d-c81a-bca2-06e6056229ae','jedediah60@example.org','$2y$10$jYEqZ9A/gDoQz0BxhKcSdessFU0mg4J1kTfAPXCXrFh6tPXCM5nqu','https://upload.wikimedia.org/wikipedia/commons/thumb/1/12/User_icon_2.svg/640px-User_icon_2.svg.png',NULL,'client',0),(5,'2021-12-21 16:57:56','2021-12-21 16:57:56','Miss Maritza Buckridge','Kenya Gislason IV','6da95e58-9a1d-119d-c966-e81f9383b172','brandon.haley@example.net','$2y$10$hbdK8nca0ISmRrWYoxgDoOh9EvC93ZwXzQMABj9Z.AH.wn1Jxwcwq','https://upload.wikimedia.org/wikipedia/commons/thumb/1/12/User_icon_2.svg/640px-User_icon_2.svg.png',NULL,'client',0),(42,'2022-01-11 17:55:20','2022-01-24 14:56:22','Adrian','Ramos','b7fe03d0-3846-6884-a022-bc486d4ef00a','ramos.urenya.adrian@gmail.com','$2y$10$jd00vA8t3bIVXd1ewuiwhOhDjc95ZymZ66gCHhdIRy6e1n.1g0Jsm','https://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/User_font_awesome.svg/2048px-User_font_awesome.svg.png',NULL,'admin',0),(43,'2022-01-12 19:11:30','2022-01-12 19:15:20','Ejemplo','Ejemplo','85c574cf-5050-ee62-9f6d-d4aab566b4d0','ejemplo@ejemplo.com','$2y$10$bFBrkdGJ/epl8me9uwdVC.KWDcUk2osxwY6q1vR5z/4OkJAUzpBUS','https://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/User_font_awesome.svg/2048px-User_font_awesome.svg.png', NULL,'admin',0), (6,'2022-01-12 19:11:30','2022-01-12 19:15:20','Yolanda','','85c574cf-5051-ee62-9f6d-d4aab566b4d0','yolanda@poliserva.com','$2a$12$s.M/zYLq4nAIGrYZLYLds.SI4rjo6E38NH7iP5OV5IZXWw87KpuRG','https://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/User_font_awesome.svg/2048px-User_font_awesome.svg.png', null,'admin',0), (7,'2022-01-12 19:11:30','2022-01-12 19:15:20','Pepe','','85c574cf-5052-ee62-9f6d-d4aab566b4d0','pepe@poliserva.com','$2a$12$2REZgo2VggzGD41MkdkK7.Z0Lpi/5RKKRjKYM1f1dpEkWOB2hKzAK','https://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/User_font_awesome.svg/2048px-User_font_awesome.svg.png', null,'admin',0), (8,'2022-01-12 19:11:30','2022-01-12 19:15:20','Javi','','85c574cf-5053-ee62-9f6d-d4aab566b4d0','javi@poliserva.com','$2a$12$5k.csEvZOq0neVhZCPBu3udyfaTI7zVtSS9UawcSLLT/3cSC7.cSC','https://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/User_font_awesome.svg/2048px-User_font_awesome.svg.png', null,'admin',0), (9,'2022-01-12 19:11:30','2022-01-12 19:15:20','Lorena','','85c574cf-5054-ee62-9f6d-d4aab566b4d0','lorena@poliserva.com','$2a$12$RxrlVcY9qXtyheBy9AxHV.8H.C93e7BXO2ThVXtFq5UuRsNRkjQCq','https://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/User_font_awesome.svg/2048px-User_font_awesome.svg.png', null,'admin',0);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-28 20:10:24
