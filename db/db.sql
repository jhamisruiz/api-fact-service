-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 22-11-2024 a las 14:17:36
-- Versión del servidor: 10.4.32-MariaDB
-- Versión de PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `facturador`
--
CREATE DATABASE IF NOT EXISTS `facturador` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `facturador`;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `usuarios`
--

DROP TABLE IF EXISTS `usuarios`;
CREATE TABLE `usuarios` (
                            `id` int(11) NOT NULL,
                            `name` varchar(50) NOT NULL,
                            `last_name` varchar(50) NOT NULL,
                            `habilitado` bit(1) NOT NULL DEFAULT b'1',
                            `email` varchar(50) NOT NULL,
                            `telefono` varchar(11) NOT NULL,
                            `user_name` varchar(30) NOT NULL,
                            `password` varchar(11) NOT NULL DEFAULT '123'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `usuarios`
--

INSERT INTO `usuarios` (`id`, `name`, `last_name`, `habilitado`, `email`, `telefono`, `user_name`, `password`) VALUES
                                                                                                                   (1, 'jhoel', 'ruiz', b'1', 'jhamis@gmail.com', '987456321', 'jruiz', '123'),
                                                                                                                   (2, 'jhamis', 'castillo', b'1', 'jcastillo@gmail.com', '987456331', 'jcastillo', '123');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `usuarios`
--
ALTER TABLE `usuarios`
    ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `usuarios`
--
ALTER TABLE `usuarios`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

