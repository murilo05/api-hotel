-- --------------------------------------------------------
-- Servidor:                     127.0.0.1
-- Versão do servidor:           10.4.25-MariaDB - Source distribution
-- OS do Servidor:               Linux
-- HeidiSQL Versão:              12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Copiando estrutura do banco de dados para tc
CREATE DATABASE IF NOT EXISTS `tc` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `tc`;

-- Copiando estrutura para tabela tc.agendamentos
CREATE TABLE IF NOT EXISTS `agendamentos` (
  `cnpj` varchar(50) DEFAULT NULL,
  `horario_inicial` time DEFAULT NULL,
  `horario_final` time DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Copiando dados para a tabela tc.agendamentos: ~0 rows (aproximadamente)

-- Copiando estrutura para tabela tc.horarios
CREATE TABLE IF NOT EXISTS `horarios` (
  `horarios_iniciais` time DEFAULT NULL,
  `horarios_finais` time DEFAULT NULL,
  `disponibilidade` tinyint(4) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Copiando dados para a tabela tc.horarios: ~11 rows (aproximadamente)
INSERT INTO `horarios` (`horarios_iniciais`, `horarios_finais`, `disponibilidade`) VALUES
	('08:00:00', '09:00:00', 1),
	('09:00:00', '10:00:00', 1),
	('10:00:00', '11:00:00', 1),
	('11:00:00', '12:00:00', 1),
	('12:00:00', '13:00:00', 1),
	('13:00:00', '14:00:00', 1),
	('14:00:00', '15:00:00', 1),
	('15:00:00', '16:00:00', 1),
	('16:00:00', '17:00:00', 1),
	('17:00:00', '18:00:00', 1),
	('18:00:00', '19:00:00', 1);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;