// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

// ###############################$
// #########################&%$###$
// #################@%%##&|!|&####$
// ##############$|!!%$|!!!|@#####$
// ###########%!!!!!!!!;;!%#######$
// ########%!!!;;;!!;;;;!$########$
// #####&|!!;;;;;;;;;;!!||!$######$
// ###&|!!;;;;;;;;;;;!!!!!&#######$
// ##&!!;;;;;;;;;;;;;;;!|@########$
// ##%!;;;:'``':;;;;;;!%@#########$
// ##$!;;;:`..`:;;;;!!$###########$
// ##@|!;;;::::;;;!!|@############$
// ####%!!;;;;;;!!|&##############$
// ######&%!!!!|%@################$
// ###############################$
// LOGO IMAGE : https://avatars0.githubusercontent.com/u/48612456?s=200&v=4

package app

// MQServer is message queue.
type MQServer struct {
	// Bind IP
	Address string
	// Bind Port 0-65535
	Port uint16
	// Allow access IP
	AllowIP []string
	// Auth password
	Password string
}
