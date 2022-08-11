-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Versi server:                 10.4.24-MariaDB - mariadb.org binary distribution
-- OS Server:                    Win64
-- HeidiSQL Versi:               12.0.0.6468
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- membuang struktur untuk table enterprise022.accounts
CREATE TABLE IF NOT EXISTS `accounts` (
  `id` bigint(20) DEFAULT NULL,
  `accountcode` varchar(50) DEFAULT NULL,
  `parentaccountid` int(11) DEFAULT NULL,
  `accountname` varchar(50) DEFAULT NULL,
  `currencyid` varchar(50) DEFAULT NULL,
  `isdebit` int(11) DEFAULT NULL,
  `accounttype` int(11) DEFAULT NULL,
  `isdisabled` int(11) DEFAULT NULL,
  `requirecostcenter` int(11) DEFAULT NULL,
  `allowallcostcenters` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Membuang data untuk tabel enterprise022.accounts: ~372 rows (lebih kurang)
INSERT INTO `accounts` (`id`, `accountcode`, `parentaccountid`, `accountname`, `currencyid`, `isdebit`, `accounttype`, `isdisabled`, `requirecostcenter`, `allowallcostcenters`) VALUES
	(1, '1', 0, 'Aktiva', '', 1, 1, 0, 0, 1),
	(2, '11', 1, 'Aktiva Lancar', '', 1, 1, 0, 0, 1),
	(3, '1110', 2, 'Kas', '', 1, 1, 0, 0, 1),
	(4, '1110.001', 3, 'Kas Kecil IDR Jakarta', '', 1, 3, 0, 0, 1),
	(5, '1110.002', 3, 'Kas Kecil USD Jakarta', '', 1, 3, 0, 0, 1),
	(6, '1110.003', 3, 'Kas Kecil Surabaya', '', 1, 3, 0, 0, 1),
	(7, '1110.004', 3, 'Kas Kecil IDR Medan', '', 1, 3, 0, 0, 1),
	(8, '1110.005', 3, 'Kas Kecil USD Medan', '', 1, 3, 0, 0, 1),
	(9, '1110.006', 3, 'Kas Jasa', '', 1, 3, 0, 0, 1),
	(10, '1110.007', 3, 'Kas Trading', '', 1, 3, 0, 0, 1),
	(11, '1110.008', 3, 'Cash Temporary', '', 1, 3, 0, 0, 1),
	(12, '1110.009', 3, 'Kas Kecil KIM', '', 1, 3, 0, 0, 1),
	(13, '1110.010', 3, 'Kas Mustafa (Medan)', '', 1, 3, 0, 0, 1),
	(15, '1120', 2, 'BANK', '', 1, 1, 0, 0, 1),
	(16, '1120.001', 15, 'BCA - 7160165558 (IDR)', '', 1, 3, 0, 0, 1),
	(17, '1120.002', 15, 'BCA - Indah- 022140038 (USD)', '', 1, 3, 0, 0, 1),
	(18, '1120.003', 15, 'BCA - Rusman - 8280070808 (IDR)', '', 1, 3, 0, 0, 1),
	(19, '1120.004', 15, 'HSBC- 918 043324-074 (IDR)', '', 1, 3, 0, 0, 1),
	(20, '1120.005', 15, 'HSBC - 800 107492-117 (USD)', '', 1, 3, 0, 0, 1),
	(21, '1120.006', 15, 'DANAMON - 003 6131563 (IDR)', '', 1, 3, 0, 0, 1),
	(22, '1120.007', 15, 'MAYBANK - 2.208.551696', '', 1, 3, 0, 0, 1),
	(23, '1120.008', 15, 'BCA - 0222511851 (Surabaya)', '', 1, 3, 0, 0, 1),
	(24, '1120.009', 15, 'BCA - 0222426888 (USD)', '', 1, 3, 0, 0, 1),
	(25, '1120.010', 15, 'BCA - 0222416688 (IDR)', '', 1, 3, 0, 0, 1),
	(26, '1120.011', 15, 'HSBC - 800-065872-075 (IDR) (DML)', '', 1, 3, 0, 0, 1),
	(27, '1120.012', 15, 'HSBC - 800-065872-117 (USD) (DML)', '', 1, 3, 0, 0, 1),
	(28, '1120.013', 15, 'HSBC - 918.043324-117 (USD)', '', 1, 3, 0, 0, 1),
	(29, '1120.014', 15, 'HSBC - 918.043324-900 (IDR)', '', 1, 3, 0, 0, 1),
	(30, '1120.015', 15, 'DEPOSITO (6688 IDR & 6888 USD)', '', 1, 3, 0, 0, 1),
	(31, '1120.016', 15, 'BCA - 806-0303810 (IDR)', '', 1, 3, 0, 0, 1),
	(32, '1120.017', 15, 'BCA - 0224270808 (IDR - RH)-pulsa', '', 1, 3, 0, 0, 1),
	(33, '1120.018', 15, 'HSBC - 918.043324-901 (IDR)', '', 1, 3, 0, 0, 1),
	(34, '1120.019', 15, 'HSBC - 918 043324-075 (IDR)', '', 1, 3, 0, 0, 1),
	(35, '1120.020', 15, 'MDR - 1680001650850 (IDR) (Master)', '', 1, 3, 0, 0, 1),
	(36, '1120.021', 15, 'MDR - 1680091650851 (IDR) (JKT)', '', 1, 3, 0, 0, 1),
	(37, '1120.022', 15, 'MDR - 1680019690856 (IDR) (Pengurus)', '', 1, 3, 0, 0, 1),
	(38, '1120.023', 15, 'MDR - 1680081650853 (IDR) (SBY)', '', 1, 3, 0, 0, 1),
	(39, '1120.024', 15, 'MDR - 1680012340855 (IDR) (MDN)', '', 1, 3, 0, 0, 1),
	(40, '1120.025', 15, 'MDR - 1680034560852 (IDR) (Pengurus)', '', 1, 3, 0, 0, 1),
	(41, '1120.026', 15, 'BCA - 1958688899 IDR PT (Kredit)', '', 1, 3, 0, 0, 1),
	(42, '1120.027', 15, 'BCA - 1951959885 PT Medan', '', 1, 3, 0, 0, 1),
	(43, '1120.028', 15, 'BCA - 1958098889 PT', '', 1, 3, 0, 0, 1),
	(44, '1120.029', 15, 'BCA - 806-0303801 PT (IDR)', '', 1, 3, 0, 0, 1),
	(45, '1130', 2, 'Persediaan Barang Dagang', '', 1, 1, 0, 0, 1),
	(46, '1130.001', 45, 'Persediaan Flexibag', '', 1, 3, 0, 0, 1),
	(47, '1130.002', 45, 'Persediaan Oleo Chemicals', '', 1, 3, 0, 0, 1),
	(48, '1130.003', 45, 'Persediaan Umum', '', 1, 3, 0, 0, 1),
	(51, '1140', 2, 'Persediaan Bahan Pembantu', '', 1, 1, 0, 0, 1),
	(52, '1140.001', 51, 'Persediaan Barang Pendukung Besi', '', 1, 3, 0, 0, 1),
	(53, '1140.002', 51, 'Persediaan Barang Pendukung Single Face', '', 1, 3, 0, 0, 1),
	(54, '1140.003', 51, 'Persediaan Barang Pendukung Bulkhead', '', 1, 3, 0, 0, 1),
	(55, '1140.004', 51, 'Persediaan Bahan Pembantu PS Foam', '', 1, 3, 0, 0, 1),
	(56, '1150', 2, 'Piutang Usaha', '', 1, 1, 0, 0, 1),
	(57, '1150.001', 56, 'Piutang Dagang IDR', '', 1, 3, 0, 0, 1),
	(58, '1150.002', 56, 'Piutang Dagang USD', '', 1, 3, 0, 0, 1),
	(59, '1150.003', 56, 'Piutang atas Jaminan', '', 1, 3, 0, 0, 1),
	(60, '1150.004', 56, 'Piutang Atas Jaminan berupa Giro', '', 1, 3, 0, 0, 1),
	(61, '1150.005', 56, 'Piutang Atas PPh 23 Supplier', '', 1, 3, 0, 0, 1),
	(62, '1160', 2, 'Piutang Non Usaha', '', 1, 1, 0, 0, 1),
	(63, '1160.001', 62, 'Piutang Karyawan', '', 1, 3, 0, 0, 1),
	(65, '1160.003', 62, 'Piutang Direksi', '', 1, 3, 0, 0, 1),
	(66, '1160.004', 62, 'Kas Bon Operasional', '', 1, 3, 0, 0, 1),
	(67, '1170', 2, 'PIUTANG ANTAR UNIT', '', 1, 1, 0, 0, 1),
	(68, '1170.001', 67, 'Piutang Unit Medan', '', 1, 3, 0, 0, 1),
	(69, '1170.002', 67, 'Piutang Unit Jakarta', '', 1, 3, 0, 0, 1),
	(70, '1170.003', 67, 'Piutang Unit Surabaya', '', 1, 3, 0, 0, 1),
	(71, '1180', 2, 'Piutang Hubungan Istimewa', '', 1, 1, 0, 0, 1),
	(72, '12', 1, 'Uang Muka Pajak', '', 0, 1, 0, 0, 1),
	(73, '1200', 72, 'Uang Muka Pajak', '', 0, 1, 0, 0, 1),
	(74, '1200.001', 73, 'PPh Ps 21 karyawan', '', 1, 3, 0, 0, 1),
	(75, '1200.002', 73, 'PPh Ps 22', '', 1, 3, 0, 0, 1),
	(76, '1200.003', 73, 'PPh Ps 23 ( kredit pajak )', '', 1, 3, 0, 0, 1),
	(77, '1200.004', 73, 'Angsuran PPh Ps 25', '', 1, 3, 0, 0, 1),
	(78, '1200.005', 73, 'PPN Masukan', '', 1, 3, 0, 0, 1),
	(79, '1200.006', 73, 'PPN Masukan Import', '', 1, 3, 0, 0, 1),
	(80, '1210', 72, 'Biaya Dibayar Dimuka', '', 1, 1, 0, 0, 1),
	(81, '1210.001', 80, 'Biaya Dibayar Dimuka Gedung', '', 1, 3, 0, 0, 1),
	(82, '1210.002', 80, 'Biaya Dibayar Dimuka Kantor', '', 1, 3, 0, 0, 1),
	(83, '1220', 72, 'Asuransi Dibayar Dimuka', '', 1, 1, 0, 0, 1),
	(84, '1220.001', 83, 'Asuransi Dibayar Dimuka - Gedung/banguna', '', 1, 3, 0, 0, 1),
	(85, '1220.002', 83, 'Asuransi Dibayar Dimuka - Kendaraan', '', 1, 3, 0, 0, 1),
	(86, '1230', 72, 'Uang Muka Pembelian', '', 1, 1, 0, 0, 1),
	(87, '1230.001', 86, 'Uang Muka Pembelian', '', 1, 3, 0, 0, 1),
	(88, '1240', 72, 'Klaim Dibayar Dimuka', '', 1, 1, 0, 0, 1),
	(89, '1240.001', 88, 'Klaim Dibayar Dimuka', '', 1, 3, 0, 0, 1),
	(90, '13', 1, 'Penyertaan', '', 1, 1, 0, 0, 1),
	(91, '1300', 90, 'Penyertaan', '', 1, 1, 0, 0, 1),
	(92, '1300.001', 91, 'Penyertaan PT. Cimi Logistic', '', 1, 3, 0, 0, 1),
	(93, '14', 1, 'Aktiva Tetap', '', 1, 1, 0, 0, 1),
	(94, '1410', 93, 'Aktiva Tetap', '', 1, 1, 0, 0, 1),
	(95, '1410.001', 94, 'Tanah dan Bangunan', '', 1, 3, 0, 0, 1),
	(96, '1410.002', 94, 'Gedung', '', 1, 3, 1, 0, 1),
	(97, '1410.003', 94, 'Kendaraan', '', 1, 3, 0, 0, 1),
	(98, '1410.004', 94, 'Peralatan Kantor', '', 1, 3, 0, 0, 1),
	(99, '1410.005', 94, 'Alat Alat Lapangan', '', 1, 3, 0, 0, 1),
	(100, '1420', 93, 'Akumulasi Penyusutan', '', 1, 1, 0, 0, 1),
	(101, '1420.001', 100, 'Ak Peny Tanah & Bangunan', '', 1, 3, 0, 0, 1),
	(102, '1420.002', 100, 'Ak Peny Kendaraan', '', 1, 3, 0, 0, 1),
	(103, '1420.003', 100, 'Ak Peny Peralatan Kantor', '', 1, 3, 0, 0, 1),
	(104, '1420.004', 100, 'Ak Peny Alat Alat Lapangan', '', 1, 3, 0, 0, 1),
	(105, '2', 0, 'Kewajiban', '', 0, 1, 0, 0, 1),
	(106, '21', 105, 'Hutang & Ekuitas', '', 0, 1, 0, 0, 1),
	(107, '2110', 106, 'Hutang', '', 0, 1, 0, 0, 1),
	(108, '2110.001', 107, 'Hutang Usaha', '', 0, 3, 0, 0, 1),
	(109, '2110.002', 107, 'Hutang Usaha Sementara', '', 0, 3, 0, 0, 1),
	(110, '2110.003', 107, 'Acrual Biaya', '', 0, 3, 0, 0, 1),
	(111, '2120', 106, 'Hutang Pajak', '', 0, 1, 0, 0, 1),
	(112, '2120.001', 111, 'PPh Ps 21 karyawan terhutang', '', 0, 3, 0, 0, 1),
	(113, '2120.002', 111, 'PPh Ps 23 yang dipotong', '', 0, 3, 0, 0, 1),
	(114, '2120.003', 111, 'PPh Ps 4 (2) atas pajak final', '', 0, 3, 0, 0, 1),
	(115, '2120.004', 111, 'PPh Ps 25 terhutang', '', 0, 3, 0, 0, 1),
	(116, '2120.005', 111, 'PPh Ps 29 tahunan', '', 0, 3, 0, 0, 1),
	(117, '2120.006', 111, 'PPN Keluaran', '', 0, 3, 0, 0, 1),
	(118, '2120.007', 111, 'PPh 21 Vendor', '', 0, 3, 0, 0, 1),
	(119, '2120.008', 111, 'Hutang PPh 22', '', 0, 3, 0, 0, 1),
	(120, '2130', 106, 'Hutang Jangka Panjang', '', 0, 1, 0, 0, 1),
	(121, '2130.001', 120, 'Hutang Kepada Pemegang Saham', '', 0, 3, 0, 0, 1),
	(122, '2130.002', 120, 'Hutang Leasing', '', 0, 3, 0, 0, 1),
	(123, '2130.003', 120, 'Bi.Gaji Yang Masih Harus Dibayarkan', '', 0, 3, 0, 0, 1),
	(124, '2140', 106, 'Uang Muka Penjualan', '', 0, 1, 0, 0, 1),
	(125, '2140.001', 124, 'Uang Muka /DP Penjualan', '', 0, 3, 0, 0, 1),
	(126, '2150', 106, 'Jaminan', '', 0, 1, 0, 0, 1),
	(127, '2150.001', 126, 'Jaminan', '', 0, 3, 0, 0, 1),
	(128, '2150.002', 126, 'Jaminan berupa Giro', '', 0, 3, 0, 0, 1),
	(129, '2160', 106, 'Titipan Direksi', '', 0, 1, 0, 0, 1),
	(130, '2160.001', 129, 'Titipan Direksi Agri Oils', '', 0, 3, 0, 0, 1),
	(131, '3', 0, 'MODAL & EKUITAS', '', 0, 1, 0, 0, 1),
	(132, '31', 131, 'MODAL', '', 0, 1, 0, 0, 1),
	(133, '3110', 132, 'MODAL', '', 0, 1, 0, 0, 1),
	(134, '3110.001', 133, 'Modal Disetor', '', 0, 3, 0, 0, 1),
	(135, '32', 131, 'LABA/RUGI DITAHAN', '', 0, 1, 0, 0, 1),
	(136, '3210', 135, 'LABA /RUGI DITAHAN', '', 0, 1, 0, 0, 1),
	(137, '3210.001', 136, 'Laba Rugi Tahun Berjalan', '', 0, 3, 0, 0, 1),
	(138, '3210.002', 136, 'Laba Rugi Bulan Berjalan', '', 0, 3, 0, 0, 1),
	(139, '3210.003', 136, 'Laba Rugi Tahun Lalu', '', 0, 3, 0, 0, 1),
	(140, '3210.004', 136, 'Penalty', '', 0, 3, 0, 0, 1),
	(141, '3210.005', 136, 'Equitas Saldo Awal', '', 0, 3, 0, 0, 1),
	(142, '33', 131, 'PRIVE', '', 0, 1, 0, 0, 1),
	(143, '3310', 142, 'PRIVE', '', 0, 1, 0, 0, 1),
	(144, '3310.001', 143, 'Prive Direksi', '', 0, 3, 0, 0, 1),
	(145, '4', 0, 'Pendapatan', '', 0, 2, 0, 0, 1),
	(146, '41', 145, 'Pendapatan', '', 0, 2, 0, 0, 1),
	(147, '4110', 146, 'Pendapatan Trading', '', 0, 2, 0, 0, 1),
	(148, '4110.001', 147, 'Penjualan Barang Dagang', '', 0, 4, 0, 0, 1),
	(149, '4110.002', 147, 'Retur Penjualan', '', 0, 4, 0, 0, 1),
	(150, '4120', 146, 'Pendapatan Jasa', '', 0, 2, 0, 0, 1),
	(151, '4120.001', 150, 'Pendapatan Jasa Trucking', '', 0, 4, 0, 0, 1),
	(152, '4120.002', 150, 'Pendapatan Komisi NCS', '', 0, 4, 0, 0, 1),
	(153, '4120.003', 150, 'Pendapatan Komisi EWAY', '', 0, 4, 0, 0, 1),
	(154, '4120.004', 150, 'Pendapatan Komisi Bulk Global', '', 0, 4, 0, 0, 1),
	(155, '4120.005', 150, 'Pendapatan Dari Principal', '', 0, 4, 0, 0, 1),
	(156, '4120.006', 150, 'Pendapatan Komisi ATI', '', 0, 4, 0, 0, 1),
	(157, '4120.007', 150, 'Pendapatan komisi HengCheng', '', 0, 4, 0, 0, 1),
	(158, '4120.008', 150, 'Pendapatan komisi Guang Zhou', '', 0, 4, 0, 0, 1),
	(159, '4120.009', 150, 'Pendapatan komisi DJD', '', 0, 4, 0, 0, 1),
	(160, '4120.010', 150, 'Pendapatan Komisi Goodrich', '', 0, 4, 0, 0, 1),
	(161, '4120.011', 150, 'Pendapatan Reimbursement', '', 0, 4, 1, 0, 1),
	(162, '4120.012', 150, 'Pendapatan Jasa Lainnya', '', 0, 4, 0, 0, 1),
	(163, '4120.013', 150, 'Pendapatan Komisi Lainnya', '', 0, 4, 0, 0, 1),
	(164, '4130', 146, 'Biaya Pembelian', '', 1, 2, 0, 0, 1),
	(165, '4130.001', 164, 'Bi. Pembelian', '', 1, 4, 0, 0, 1),
	(166, '4140', 146, 'HPP Barang Dagang', '', 1, 2, 0, 0, 1),
	(167, '4140.001', 166, 'HPP Packaging', '', 1, 4, 0, 0, 1),
	(168, '4140.002', 166, 'HPP Oleo Chemicals', '', 1, 4, 0, 0, 1),
	(169, '4140.003', 166, 'HPP Palm Kernel Shell', '', 1, 4, 1, 0, 1),
	(170, '4140.004', 166, 'Penyesuaian Barang Trading', '', 1, 4, 0, 0, 1),
	(171, '4150', 146, 'HPP Barang Pendukung', '', 1, 2, 0, 0, 1),
	(172, '4150.001', 171, '"HPP Besi', ' Metal Bars"', 0, 1, 4, 0, 0),
	(173, '4150.002', 171, 'HPP Single Face', '', 1, 4, 0, 0, 1),
	(174, '4150.003', 171, 'HPP Bulkhead', '', 1, 4, 0, 0, 1),
	(175, '4150.004', 171, 'HPP PS FOAM', '', 1, 4, 0, 0, 1),
	(176, '4150.005', 171, 'HPP Heating Pad', '', 1, 4, 0, 0, 1),
	(177, '5', 0, 'Biaya - Biaya', '', 1, 2, 0, 0, 1),
	(178, '51', 177, 'Biaya Operasional', '', 1, 2, 0, 0, 1),
	(179, '5110', 178, 'Biaya Operasional Langsung', '', 1, 2, 0, 0, 1),
	(180, '5110.001', 179, 'Bi. Freight', '', 1, 4, 0, 0, 1),
	(181, '5110.002', 179, 'Bi. Penumpukan', '', 1, 4, 0, 0, 1),
	(182, '5110.003', 179, 'Bi. Lift On and Off', '', 1, 4, 0, 0, 1),
	(183, '5110.005', 179, 'Bi. Handling Lokal', '', 1, 4, 0, 0, 1),
	(184, '5110.006', 179, 'Bi. Handling Export', '', 1, 4, 0, 0, 1),
	(185, '5110.007', 179, 'Bi. Handling Import', '', 1, 4, 0, 0, 1),
	(186, '5110.008', 179, 'Bi. PEB dan PIB', '', 1, 4, 0, 0, 1),
	(187, '5110.009', 179, 'Bi. Uang Jalan Trucking', '', 1, 4, 0, 0, 1),
	(188, '5110.010', 179, 'Bi. Dokumen', '', 1, 4, 0, 0, 1),
	(189, '5110.011', 179, 'Bi. Sparepart Trado', '', 1, 4, 0, 0, 1),
	(190, '5110.012', 179, 'Bi. DG', '', 1, 4, 0, 0, 1),
	(191, '5110.013', 179, 'Bi. Pelancar', '', 1, 4, 0, 0, 1),
	(192, '5110.014', 179, 'Bi. Demurage', '', 1, 4, 0, 0, 1),
	(193, '5110.015', 179, 'Bi. Jalur Merah', '', 1, 4, 0, 0, 1),
	(194, '5110.016', 179, 'Bi. Jalur Kuning', '', 1, 4, 0, 0, 1),
	(195, '5110.017', 179, 'Bi. Seal', '', 1, 4, 0, 0, 1),
	(196, '5110.018', 179, 'Bi. Gaji Freelance', '', 1, 4, 0, 0, 1),
	(197, '5110.019', 179, 'Bi. Keamanan Trado', '', 1, 4, 0, 0, 1),
	(198, '5110.020', 179, 'Bi. keperluan Isotank', '', 1, 4, 0, 0, 1),
	(199, '5110.021', 179, 'Bi. Pembelian Ban Trado', '', 1, 4, 0, 0, 1),
	(200, '5110.022', 179, 'Bi. Depot Kontener', '', 1, 4, 0, 0, 1),
	(201, '5110.023', 179, 'Bi. STNK / KIR', '', 1, 4, 0, 0, 1),
	(202, '5110.024', 179, 'Bi. Sewa Trucking', '', 1, 4, 0, 0, 1),
	(203, '5110.025', 179, 'Bi. Inap Trucking', '', 1, 4, 0, 0, 1),
	(204, '5110.026', 179, 'Bi. Stand By Sopir', '', 1, 4, 0, 0, 1),
	(205, '5110.027', 179, 'Bi. Nginap Supir', '', 1, 4, 0, 0, 1),
	(206, '5110.028', 179, 'Bi. Perlengkapan Safety Kendaraan', '', 1, 4, 0, 0, 1),
	(207, '5110.029', 179, 'Bi. Lain-lain Operasional Langsung', '', 1, 4, 0, 0, 1),
	(208, '5110.030', 179, 'Bi. Sewa Chassis RH', '', 1, 4, 0, 0, 1),
	(209, '5110.031', 179, 'Bi. Reimbursement', '', 1, 4, 1, 0, 1),
	(210, '5110.032', 179, 'Bi. Pemeliharaan & Perbaikan Forklift', '', 1, 4, 0, 0, 1),
	(211, '5110.033', 179, 'Bi. Tagihan Tahunan GPS Trado', '', 1, 4, 0, 0, 1),
	(212, '5110.034', 179, 'Low Value Asset', '', 1, 4, 0, 0, 1),
	(213, '5110.035', 179, 'Bi. Storage', '', 1, 4, 0, 0, 1),
	(214, '5110.036', 179, 'Bi. Cleaning', '', 1, 4, 0, 0, 1),
	(215, '5110.037', 179, 'Bi. Survey', '', 1, 4, 0, 0, 1),
	(216, '5110.038', 179, 'Bi. Pemeliharaan Gandengan', '', 1, 4, 0, 0, 1),
	(217, '5110.039', 179, 'Bi. Kerugian Operasional', '', 1, 4, 0, 0, 1),
	(218, '5110.040', 179, 'Bi. Tank Fee', '', 1, 4, 0, 0, 1),
	(219, '5110.041', 179, 'Bi. Teknisi', '', 1, 4, 0, 0, 1),
	(220, '5110.042', 179, 'Bi. Pemeliharaan Trado', '', 1, 4, 0, 0, 1),
	(221, '5110.043', 179, 'Bi. Pembelian Ban Mobil Box', '', 1, 4, 0, 0, 1),
	(222, '5110.044', 179, 'Bi. Pemeliharaan Mobil Box', '', 1, 4, 0, 0, 1),
	(223, '5110.045', 179, 'Bi. Operasional Supir Mitra', '', 1, 4, 0, 0, 1),
	(224, '5110.046', 179, 'Bi. Repair', '', 1, 4, 0, 0, 1),
	(225, '5110.047', 179, 'Bi. PIB', '', 1, 4, 0, 0, 1),
	(227, '5110.048', 179, 'Bi. Uang Jalan utk Box/Trading', '', 1, 4, 0, 0, 1),
	(228, '5110.049', 179, 'Bi. Asuransi Penjualan', '', 1, 4, 0, 0, 1),
	(229, '5110.050', 179, 'Bi. Pulsa/Voucher', '', 1, 4, 0, 0, 1),
	(230, '5110.051', 179, 'Bi. BBM kendaraan', '', 1, 4, 0, 0, 1),
	(231, '5110.052', 179, 'Bi. Parkir dan Tol', '', 1, 4, 0, 0, 1),
	(232, '5110.053', 179, 'Bi. SPSI / Bongkar muat', '', 1, 4, 0, 0, 1),
	(233, '5110.054', 179, 'Bi. Taktis & Koordinasi', '', 1, 4, 0, 0, 1),
	(234, '5110.055', 179, 'Bi. Pindah Antar Gudang', '', 1, 4, 0, 0, 1),
	(235, '5110.056', 179, 'Bi. Kerusakan Flexibag / Tank', '', 1, 4, 0, 0, 1),
	(236, '5110.057', 179, 'Bi. Kerusakan Barang Trading', '', 1, 4, 0, 0, 1),
	(237, '5110.058', 179, 'Bi. Pemeliharaan/Perbaikan Kalmar', '', 1, 4, 0, 0, 1),
	(238, '5110.059', 179, 'Bi. Penyusutan Barang Kimia', '', 1, 4, 0, 0, 1),
	(239, '5110.060', 179, 'Bi. Klaim Susut', '', 1, 4, 0, 0, 1),
	(240, '5110.061', 179, 'Bi. Lain-lain Penjualan Trading', '', 1, 4, 0, 0, 1),
	(241, '5110.062', 179, 'Bi. Kawalan Trucking', '', 1, 4, 0, 0, 1),
	(242, '5110.063', 179, 'Bi. Gaji Karyawan Operasional', '', 1, 4, 0, 0, 1),
	(243, '5110.064', 179, 'Bi. Repo', '', 1, 4, 0, 0, 1),
	(244, '5110.065', 179, 'Bi. Keperluan Container', '', 1, 4, 0, 0, 1),
	(245, '5110.066', 179, 'Bi. Pengiriman Barang', '', 1, 4, 0, 0, 1),
	(246, '5110.067', 179, 'Bi. Komisi Penjualan', '', 1, 4, 0, 0, 1),
	(247, '5110.068', 179, 'Bi. Perjalanan Dinas Operasional', '', 1, 4, 0, 0, 1),
	(248, '52', 177, 'BIAYA ADMINISTRASI & UMUM', '', 1, 2, 0, 0, 1),
	(249, '5210', 248, 'BIAYA ADMINISTRASI & UMUM', '', 1, 2, 0, 0, 1),
	(250, '5210.001', 249, 'Bi. Gaji Karyawan', '', 1, 4, 0, 0, 1),
	(251, '5210.002', 249, 'Bi. THR/Bonus/Parcel', '', 1, 4, 0, 0, 1),
	(252, '5210.003', 249, 'Bi. Lembur', '', 1, 4, 0, 0, 1),
	(253, '5210.004', 249, 'Bi. Bunga Kredit Leasing', '', 1, 4, 0, 0, 1),
	(254, '5210.005', 249, 'Bi. BPJS Kesehatan', '', 1, 4, 0, 0, 1),
	(255, '5210.006', 249, 'Bi. BPJS Ketenagakerjaan', '', 1, 4, 0, 0, 1),
	(256, '5210.007', 249, 'Bi. Pemeliharaan Kendaraan Kantor', '', 1, 4, 0, 0, 1),
	(257, '5210.008', 249, 'Bi. Pemeliharaan & Perbaikan Gedung', '', 1, 4, 0, 0, 1),
	(258, '5210.009', 249, 'Bi. Pemeliharaan/Perbaikan Peralatan Ktr', '', 1, 4, 0, 0, 1),
	(259, '5210.010', 249, 'Bi. Penyusutan - Gedung/Bangunan', '', 1, 4, 0, 0, 1),
	(260, '5210.011', 249, 'Bi. Penyusutan - Kendaraan', '', 1, 4, 0, 0, 1),
	(261, '5210.012', 249, 'Bi. Penyusutan - Peralatan Kantor', '', 1, 4, 0, 0, 1),
	(262, '5210.013', 249, 'Bi. Penyusutan - Mesin / Alat Berat', '', 1, 4, 1, 0, 1),
	(263, '5210.014', 249, 'Bi. Penyusutan - Alat Lapangan', '', 1, 4, 0, 0, 1),
	(264, '5210.015', 249, 'Bi. Asuransi Umum dan Kantor', '', 1, 4, 0, 0, 1),
	(265, '5210.016', 249, 'Bi. Alat Tulis Kantor', '', 1, 4, 0, 0, 1),
	(266, '5210.017', 249, 'Bi. Perlengkapan Kantor', '', 1, 4, 0, 0, 1),
	(267, '5210.018', 249, '"Bi. Foto Copy', ' Percetakan & Penjilidan"', 0, 1, 4, 0, 0),
	(268, '5210.019', 249, 'Bi. Benda Benda Pos & Materai', '', 1, 4, 0, 0, 1),
	(269, '5210.020', 249, 'Bi. Kirim Dokumen', '', 1, 4, 0, 0, 1),
	(270, '5210.021', 249, '"Bi. Telp', ' Fax & Internet"', 0, 1, 4, 0, 0),
	(271, '5210.022', 249, 'Bi. Listrik & Air PAM', '', 1, 4, 0, 0, 1),
	(272, '5210.023', 249, 'Bi. Pulsa/Voucher', '', 1, 4, 0, 0, 1),
	(273, '5210.024', 249, 'Bi. Sewa Gudang/Lahan', '', 1, 4, 0, 0, 1),
	(274, '5210.025', 249, 'Bi. BBM Kendaraan', '', 1, 4, 0, 0, 1),
	(275, '5210.026', 249, 'Bi. Perijinan Admin Umum', '', 1, 4, 0, 0, 1),
	(276, '5210.027', 249, 'Bi. Keamanan Lingkungan Kantor', '', 1, 4, 0, 0, 1),
	(277, '5210.028', 249, 'Bi. Sumbangan & Donasi', '', 1, 4, 0, 0, 1),
	(278, '5210.029', 249, 'Bi. Perjalanan Dinas', '', 1, 4, 0, 0, 1),
	(279, '5210.030', 249, 'Bi. Denda Pajak/PPn tidak tertagih', '', 1, 4, 0, 0, 1),
	(280, '5210.031', 249, 'Bi. Entertainment', '', 1, 4, 0, 0, 1),
	(281, '5210.032', 249, 'Bi. Pantry', '', 1, 4, 0, 0, 1),
	(282, '5210.033', 249, 'Bi. Jamuan Makan', '', 1, 4, 0, 0, 1),
	(283, '5210.034', 249, 'Bi. Training / Pelatihan', '', 1, 4, 0, 0, 1),
	(284, '5210.035', 249, 'Bi. Akomodasi', '', 1, 4, 0, 0, 1),
	(285, '5210.036', 249, 'Bi. Piutang tidak tertagih', '', 1, 4, 0, 0, 1),
	(286, '5210.037', 249, '"Bi. Tenaga Ahli (Notaris', ' Pengacara)"', 0, 1, 4, 0, 0),
	(287, '5210.038', 249, 'Bi. STNK/BPKB (kantor)', '', 1, 4, 0, 0, 1),
	(288, '5210.039', 249, 'Bi. Pengiriman Dokumen/Barang', '', 1, 4, 0, 0, 1),
	(289, '5210.040', 249, 'Bi. Iklan', '', 1, 4, 0, 0, 1),
	(290, '5210.041', 249, 'Bi. Pembyrn PBB', '', 1, 4, 0, 0, 1),
	(291, '5210.042', 249, 'Bi. Sewa Kantor', '', 1, 4, 0, 0, 1),
	(292, '5210.043', 249, 'Bi. Pemeliharaan & Perbaikan Alat Lapang', '', 1, 4, 0, 0, 1),
	(293, '5210.044', 249, 'Bi. Uang Makan Karyawan', '', 1, 4, 0, 0, 1),
	(294, '5210.045', 249, 'Bi. Asuransi Jiwa', '', 1, 4, 0, 0, 1),
	(295, '5210.046', 249, 'Bi. Keperluan Safety', '', 1, 4, 0, 0, 1),
	(296, '5210.047', 249, 'Bi. Penyusutan Cangkang (Palm Kernell Shell)', '', 1, 4, 0, 0, 1),
	(297, '5210.048', 249, 'Bi. Limbah B3', '', 1, 4, 0, 0, 1),
	(298, '5210.049', 249, 'Bi. APD/Safeguard', '', 1, 4, 0, 0, 1),
	(299, '5210.050', 249, 'Bi. Gaji Mitra', '', 1, 4, 0, 0, 1),
	(300, '5210.051', 249, 'Bi. Retribusi Sampah', '', 1, 4, 0, 0, 1),
	(301, '5210.099', 249, 'Bi. Lain-Lain Admin Umum', '', 1, 4, 0, 0, 1),
	(302, '6', 0, 'PENDAPATAN / BIAYA LAIN-LAIN', '', 0, 2, 0, 0, 1),
	(303, '61', 302, 'PENDAPATAN / BIAYA LAIN-LAIN', '', 0, 2, 0, 0, 1),
	(304, '6110', 303, 'PENDAPATAN LAIN LAIN', '', 0, 2, 0, 0, 1),
	(305, '6110.001', 304, 'Pendapatan Jasa Giro', '', 0, 4, 0, 0, 1),
	(306, '6110.002', 304, 'Selisih Kurs', '', 0, 4, 0, 0, 1),
	(307, '6110.003', 304, 'Pendapatan Lain-Lain', '', 0, 4, 0, 0, 1),
	(308, '6110.004', 304, 'Bunga Deposito', '', 0, 4, 0, 0, 1),
	(309, '6110.005', 304, 'Lebih Bayar', '', 0, 4, 1, 0, 1),
	(310, '6110.006', 304, 'Pendapatan Bunga Pinjaman', '', 0, 4, 0, 0, 1),
	(311, '62', 302, 'BIAYA LAIN - LAIN', '', 1, 2, 0, 0, 1),
	(312, '6210', 311, 'BIAYA LAIN - LAIN', '', 1, 2, 0, 0, 1),
	(313, '6210.001', 312, 'Bi. Provisi dan Adm bank', '', 1, 4, 0, 0, 1),
	(314, '6210.002', 312, 'Bi. Pajak Jasa Giro', '', 1, 4, 0, 0, 1),
	(315, '6210.003', 312, 'Bi. Rounded', '', 1, 4, 0, 0, 1),
	(316, '6210.004', 312, 'Kurang Bayar', '', 1, 4, 1, 0, 1),
	(317, '6210.005', 312, 'Biaya Lain-Lain', '', 1, 4, 0, 0, 1),
	(318, '6210.006', 312, 'Bi. Bunga Pinjaman Bank', '', 1, 4, 0, 0, 1),
	(319, '6210.007', 312, 'Bi. Kasbon/Piutang Tak Tertagih', '', 1, 4, 0, 0, 1),
	(320, '6210.008', 312, 'Bi. Kerugian Penghapusan Asset', '', 1, 4, 0, 0, 1),
	(321, '6210.009', 312, 'Biaya Lain-Lain Kendaraan', '', 1, 4, 0, 0, 1),
	(322, '2110.004', 107, 'Suspend Biaya EMKL', '', 0, 3, 0, 0, 1),
	(323, '5110.069', 179, 'Bi. Materai', '', 1, 4, 0, 0, 1),
	(324, '5110.070', 179, 'Bi. Kesehatan', '', 1, 4, 0, 0, 1),
	(326, '4110.003', 147, 'Diskon Penjualan', '', 1, 4, 0, 0, 1),
	(327, '1160.002', 62, 'Lebih/Kurang Bayar', '', 1, 3, 0, 0, 1),
	(1328, '2170', 106, 'Hutang Lainnya', '', 0, 1, 0, 0, 1),
	(1329, '2170.001', 1328, 'Hutang Pengurus', '', 0, 3, 0, 0, 1),
	(1330, '2170.002', 1328, 'Hutang Teknisi', '', 0, 3, 0, 0, 1),
	(1331, '2170.003', 1328, 'Hutang Supir', '', 0, 3, 0, 0, 1),
	(1332, '5210.052', 249, 'Bi. Air Minum', '', 1, 4, 0, 0, 1),
	(1333, '1160.005', 62, 'Kas Bon Karyawan', '', 1, 3, 0, 0, 1),
	(1334, '5110.071', 179, 'Bi. Sewa Alat Berat', '', 1, 4, 0, 0, 1),
	(1335, '5210.053', 249, 'Bi. Operasional HRD', '', 1, 4, 0, 0, 1),
	(1336, '5110.072', 179, 'Bi. Service', '', 1, 4, 0, 0, 1),
	(1337, '2170.004', 1328, 'Hutang Karyawan', '', 0, 3, 0, 0, 1),
	(1338, '5110.073', 179, 'Bi. Inap Supir', '', 1, 4, 0, 0, 1),
	(1339, '5210.054', 249, 'Bi. Security', '', 1, 4, 0, 0, 1),
	(1340, '1130.006', 45, 'Persediaan In Transit', '', 1, 3, 0, 0, 1),
	(1341, '5210.055', 249, 'Bi. Keanggotaan ', '', 1, 4, 0, 0, 1),
	(1342, '4140.005', 166, 'Pemakaian Barang Trading', '', 1, 4, 0, 0, 1),
	(1343, '1200.007', 73, 'PPh Ps 4 (2)', '', 1, 3, 0, 0, 1),
	(1344, '5110.074', 179, 'Bi. Perjalanan Dinas Teknisi', '', 1, 4, 0, 0, 1),
	(1345, '1150.006', 56, 'Piutang BLT (Dana kas)', '', 1, 3, 0, 0, 1),
	(1346, '1150.007', 56, 'Piutang PT. OAK', '', 1, 3, 0, 0, 1),
	(1347, '2150.003', 126, 'Nota Jaminan', '', 0, 3, 0, 0, 1),
	(1348, '2170.005', 1328, 'Hutang Lain-Lain', '', 0, 3, 0, 0, 1),
	(1349, '5210.056', 249, 'Bi. Renovasi Garasi', '', 1, 4, 0, 0, 1),
	(1350, '1140.005', 51, 'Persediaan Heating Pad', '', 1, 3, 0, 0, 1),
	(1352, '5110.075', 179, 'Bi. Pelabuhan', '', 1, 4, 0, 0, 1),
	(1353, '4130.002', 164, 'Diskon Pembelian', '', 0, 4, 0, 0, 1),
	(1354, '5110.076', 179, 'Bi. Keperluan Trado / Trucking', '', 1, 4, 0, 0, 1),
	(1355, '5210.057', 249, 'Bi. Keperluan P3K & Obat', '', 1, 4, 0, 0, 1),
	(1356, '4110.004', 147, 'Diskon Penjualan Oleo', '', 1, 3, 1, 0, 1),
	(1357, '5210.058', 249, 'Bi. Kebutuhan Karyawan ', '', 1, 4, 0, 0, 1),
	(1358, '2120.009', 111, 'Hutang Pajak', '', 0, 3, 0, 0, 1),
	(1359, '5210.059', 249, 'Bi. Pajak', '', 1, 4, 0, 0, 1),
	(1360, '6110.007', 304, 'Pendapatan Deviden', '', 0, 4, 0, 0, 1),
	(1361, '5110.077', 179, 'Bi. Laka ', '', 1, 4, 0, 0, 1),
	(1363, '15', 1, 'Asset Lainnya', '', 1, 1, 0, 0, 1),
	(1364, '1510.001', 1363, 'Reimbursement', '', 1, 3, 0, 0, 1),
	(1365, '5210.060', 249, 'Bi. Gaji Freelance', '', 1, 4, 0, 0, 1),
	(1366, '9', 0, 'IKTHISAR L/R', '', 1, 5, 0, 0, 1),
	(1367, '1210.003', 80, 'Sewa Chasis Dibayar Dimuka', '', 1, 3, 0, 0, 1),
	(1368, '5110.078', 179, 'Bi. Keperluan Gandengan', '', 1, 4, 0, 0, 1),
	(1369, '5110.079', 179, 'Bi. Sparepart Mobil Box', '', 1, 4, 0, 0, 1),
	(1370, '1140.006', 51, 'Persediaan Accesoris Isotank', '', 1, 3, 1, 0, 1),
	(1371, '1210.004', 80, 'Biaya Dibayar Dimuka Internet/Wifi', '', 1, 3, 0, 0, 1),
	(1372, '5110.080', 179, 'Bi. Keperluan Forklift', '', 1, 4, 0, 0, 1),
	(1373, '5110.081', 179, 'Bi. Keperluan Trading', '', 1, 4, 0, 0, 1),
	(1374, '5210.061', 249, 'Bi. Tunjangan PPh 21', '', 1, 4, 0, 0, 1),
	(1375, '1130.007', 45, 'Penyesuaian Persediaan', '', 1, 3, 1, 0, 1),
	(1376, '5210.062', 249, 'Bi. Perbaikan Peralatan & Mesin', '', 1, 4, 0, 0, 1),
	(1377, '3210.006', 136, 'DEVIDEN', '', 0, 3, 0, 0, 1),
	(1378, '6210.010', 312, 'Bi. Pajak Badan', '', 1, 4, 0, 0, 1),
	(1379, '5210.063', 249, 'Bi. Instalasi Listrik/Air', '', 1, 4, 0, 0, 1),
	(1380, '5210.064', 249, 'Bi. Server Sistem ', '', 1, 4, 0, 0, 1),
	(NULL, 'ACC-TEST', 1, 'Activation', 'IDR, USD', 1, 1, 1, 1, 1),
	(NULL, 'ACC-TEST', 1, 'Activation', 'IDR, USD', 1, 1, 1, 1, 1),
	(NULL, 'ACC-TEST', 1, 'Activation', 'IDR, USD', 1, 1, 1, 1, 1);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;