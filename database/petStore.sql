-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.4.6-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             10.2.0.5599
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for petStore
DROP DATABASE IF EXISTS `petStore`;
CREATE DATABASE IF NOT EXISTS `petstore` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `petStore`;

-- Dumping structure for table petStore.pet
DROP TABLE IF EXISTS `pet`;
CREATE TABLE IF NOT EXISTS `pet` (
  `id` bigint(20) DEFAULT NULL,
  `category` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`category`)),
  `name` varchar(50) DEFAULT NULL,
  `photoUrls` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`photoUrls`)),
  `tags` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='pet';

-- Dumping data for table petStore.pet: ~8 rows (approximately)
/*!40000 ALTER TABLE `pet` DISABLE KEYS */;
INSERT IGNORE INTO `pet` (`id`, `category`, `name`, `photoUrls`, `tags`, `status`) VALUES
	(2, ' {\r\n    "id": 2,\r\n    "name": "USA"\r\n  }', 'Wang', '["c:\\\\2.jpg","c:\\\\3.jpg"]', '[{"id":3,"name":"Q"},{"id":4,"name":"C"}]', 'pending'),
	(3, ' {\r\n    "id": 3,\r\n    "name": "CHN"\r\n  }', 'CHN Cai', '["c:\\\\4.jpg","c:\\\\5.jpg"]', '[{"id":1,"name":"lovely"},{"id":2,"name":"cool"}]', 'sold'),
	(1, ' {\r\n    "id": 1,\r\n    "name": "AU"\r\n  }', 'doggie', '["c:\\\\1.jpg","c:\\\\2.jpg"]', '[{"id":1,"name":"lovely"},{"id":2,"name":"cool"}]', 'available'),
	(4, ' {\r\n    "id": 4,\r\n    "name": "JP"\r\n  }', 'JP Wang', '["c:\\\\6.jpg","c:\\\\7.jpg"]', '[{"id":5,"name":"Large"},{"id":6,"name":"Strong"}]', 'sold'),
	(5, ' {\r\n    "id": 5,\r\n    "name": "CHN"\r\n  }', 'Cai', '["c:\\\\8.jpg","c:\\\\9.jpg"]', '[{"id":7,"name":"small"},{"id":2,"name":"cool"}]', 'sold'),
	(6, '{"id":0,"name":"string"}', 'doggie', '["string"]', '[{"id":0,"name":"string"}]', 'available'),
	(7, '{"id":6,"name":"Large Dog"}', 'Big Dog', '["c:\\\\4.jpg","c:\\\\5.jpg"]', '[{"id":1,"name":"tag1"},{"id":2,"name":"tag2"}]', 'available'),
	(8, '{"id":1,"name":"Dog 8"}', 'Big Dog 2', '["c:\\\\4.jpg","c:\\\\5.jpg"]', '[{"id":1,"name":"tag1"},{"id":2,"name":"tag2"}]', 'available');
/*!40000 ALTER TABLE `pet` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
