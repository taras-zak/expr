// Code generated from Expr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package gen

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 61, 568,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3,
	41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44,
	3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 5, 46, 314, 10,
	46, 3, 47, 3, 47, 3, 47, 3, 47, 5, 47, 320, 10, 47, 3, 48, 3, 48, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 5, 48, 344,
	10, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3,
	49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3,
	49, 5, 49, 387, 10, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 50, 5, 50, 398, 10, 50, 3, 51, 3, 51, 3, 51, 7, 51, 403, 10,
	51, 12, 51, 14, 51, 406, 11, 51, 5, 51, 408, 10, 51, 3, 52, 3, 52, 3, 52,
	6, 52, 413, 10, 52, 13, 52, 14, 52, 414, 3, 52, 3, 52, 6, 52, 419, 10,
	52, 13, 52, 14, 52, 420, 5, 52, 423, 10, 52, 3, 53, 3, 53, 3, 53, 6, 53,
	428, 10, 53, 13, 53, 14, 53, 429, 3, 54, 3, 54, 7, 54, 434, 10, 54, 12,
	54, 14, 54, 437, 11, 54, 3, 55, 3, 55, 7, 55, 441, 10, 55, 12, 55, 14,
	55, 444, 11, 55, 3, 55, 3, 55, 3, 55, 7, 55, 449, 10, 55, 12, 55, 14, 55,
	452, 11, 55, 3, 55, 5, 55, 455, 10, 55, 3, 56, 6, 56, 458, 10, 56, 13,
	56, 14, 56, 459, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 7, 57, 468,
	10, 57, 12, 57, 14, 57, 471, 11, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57,
	3, 58, 3, 58, 3, 58, 3, 58, 7, 58, 482, 10, 58, 12, 58, 14, 58, 485, 11,
	58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 3, 61,
	3, 61, 3, 61, 5, 61, 499, 10, 61, 3, 62, 3, 62, 3, 62, 3, 62, 5, 62, 505,
	10, 62, 3, 63, 3, 63, 3, 63, 3, 63, 5, 63, 511, 10, 63, 3, 64, 3, 64, 5,
	64, 515, 10, 64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66,
	3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 5, 69, 534,
	10, 69, 3, 70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 5, 71, 542, 10, 71, 3,
	72, 3, 72, 5, 72, 546, 10, 72, 3, 73, 3, 73, 3, 73, 5, 73, 551, 10, 73,
	3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 7, 76, 560, 10, 76, 12,
	76, 14, 76, 563, 11, 76, 5, 76, 565, 10, 76, 3, 77, 3, 77, 3, 469, 2, 78,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59,
	31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77,
	40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95,
	49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111, 57,
	113, 58, 115, 59, 117, 60, 119, 61, 121, 2, 123, 2, 125, 2, 127, 2, 129,
	2, 131, 2, 133, 2, 135, 2, 137, 2, 139, 2, 141, 2, 143, 2, 145, 2, 147,
	2, 149, 2, 151, 2, 153, 2, 3, 2, 16, 3, 2, 51, 59, 4, 2, 50, 59, 97, 97,
	4, 2, 90, 90, 122, 122, 6, 2, 11, 11, 13, 14, 34, 34, 162, 162, 5, 2, 12,
	12, 15, 15, 8234, 8235, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 6, 2, 12,
	12, 15, 15, 41, 41, 94, 94, 11, 2, 36, 36, 41, 41, 94, 94, 100, 100, 104,
	104, 112, 112, 116, 116, 118, 118, 120, 120, 14, 2, 12, 12, 15, 15, 36,
	36, 41, 41, 50, 59, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118,
	120, 122, 122, 4, 2, 119, 119, 122, 122, 4, 2, 38, 38, 97, 97, 3, 2, 97,
	97, 4, 2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 72, 99, 104, 2, 592, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11,
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2,
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2,
	2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2,
	2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2,
	2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3,
	2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57,
	3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2,
	65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2,
	2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2,
	2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2,
	2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3,
	2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103,
	3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2,
	2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3,
	2, 2, 2, 2, 119, 3, 2, 2, 2, 3, 155, 3, 2, 2, 2, 5, 158, 3, 2, 2, 2, 7,
	169, 3, 2, 2, 2, 9, 178, 3, 2, 2, 2, 11, 187, 3, 2, 2, 2, 13, 195, 3, 2,
	2, 2, 15, 198, 3, 2, 2, 2, 17, 205, 3, 2, 2, 2, 19, 209, 3, 2, 2, 2, 21,
	213, 3, 2, 2, 2, 23, 218, 3, 2, 2, 2, 25, 222, 3, 2, 2, 2, 27, 226, 3,
	2, 2, 2, 29, 233, 3, 2, 2, 2, 31, 237, 3, 2, 2, 2, 33, 239, 3, 2, 2, 2,
	35, 241, 3, 2, 2, 2, 37, 243, 3, 2, 2, 2, 39, 245, 3, 2, 2, 2, 41, 247,
	3, 2, 2, 2, 43, 249, 3, 2, 2, 2, 45, 251, 3, 2, 2, 2, 47, 253, 3, 2, 2,
	2, 49, 255, 3, 2, 2, 2, 51, 257, 3, 2, 2, 2, 53, 259, 3, 2, 2, 2, 55, 261,
	3, 2, 2, 2, 57, 263, 3, 2, 2, 2, 59, 265, 3, 2, 2, 2, 61, 267, 3, 2, 2,
	2, 63, 271, 3, 2, 2, 2, 65, 275, 3, 2, 2, 2, 67, 277, 3, 2, 2, 2, 69, 280,
	3, 2, 2, 2, 71, 282, 3, 2, 2, 2, 73, 284, 3, 2, 2, 2, 75, 287, 3, 2, 2,
	2, 77, 290, 3, 2, 2, 2, 79, 292, 3, 2, 2, 2, 81, 294, 3, 2, 2, 2, 83, 297,
	3, 2, 2, 2, 85, 300, 3, 2, 2, 2, 87, 303, 3, 2, 2, 2, 89, 306, 3, 2, 2,
	2, 91, 313, 3, 2, 2, 2, 93, 319, 3, 2, 2, 2, 95, 343, 3, 2, 2, 2, 97, 386,
	3, 2, 2, 2, 99, 397, 3, 2, 2, 2, 101, 407, 3, 2, 2, 2, 103, 422, 3, 2,
	2, 2, 105, 424, 3, 2, 2, 2, 107, 431, 3, 2, 2, 2, 109, 454, 3, 2, 2, 2,
	111, 457, 3, 2, 2, 2, 113, 463, 3, 2, 2, 2, 115, 477, 3, 2, 2, 2, 117,
	488, 3, 2, 2, 2, 119, 492, 3, 2, 2, 2, 121, 498, 3, 2, 2, 2, 123, 504,
	3, 2, 2, 2, 125, 510, 3, 2, 2, 2, 127, 514, 3, 2, 2, 2, 129, 516, 3, 2,
	2, 2, 131, 520, 3, 2, 2, 2, 133, 526, 3, 2, 2, 2, 135, 528, 3, 2, 2, 2,
	137, 533, 3, 2, 2, 2, 139, 535, 3, 2, 2, 2, 141, 541, 3, 2, 2, 2, 143,
	545, 3, 2, 2, 2, 145, 550, 3, 2, 2, 2, 147, 552, 3, 2, 2, 2, 149, 554,
	3, 2, 2, 2, 151, 564, 3, 2, 2, 2, 153, 566, 3, 2, 2, 2, 155, 156, 7, 48,
	2, 2, 156, 157, 7, 48, 2, 2, 157, 4, 3, 2, 2, 2, 158, 159, 7, 117, 2, 2,
	159, 160, 7, 118, 2, 2, 160, 161, 7, 99, 2, 2, 161, 162, 7, 116, 2, 2,
	162, 163, 7, 118, 2, 2, 163, 164, 7, 117, 2, 2, 164, 165, 7, 89, 2, 2,
	165, 166, 7, 107, 2, 2, 166, 167, 7, 118, 2, 2, 167, 168, 7, 106, 2, 2,
	168, 6, 3, 2, 2, 2, 169, 170, 7, 103, 2, 2, 170, 171, 7, 112, 2, 2, 171,
	172, 7, 102, 2, 2, 172, 173, 7, 117, 2, 2, 173, 174, 7, 89, 2, 2, 174,
	175, 7, 107, 2, 2, 175, 176, 7, 118, 2, 2, 176, 177, 7, 106, 2, 2, 177,
	8, 3, 2, 2, 2, 178, 179, 7, 101, 2, 2, 179, 180, 7, 113, 2, 2, 180, 181,
	7, 112, 2, 2, 181, 182, 7, 118, 2, 2, 182, 183, 7, 99, 2, 2, 183, 184,
	7, 107, 2, 2, 184, 185, 7, 112, 2, 2, 185, 186, 7, 117, 2, 2, 186, 10,
	3, 2, 2, 2, 187, 188, 7, 111, 2, 2, 188, 189, 7, 99, 2, 2, 189, 190, 7,
	118, 2, 2, 190, 191, 7, 101, 2, 2, 191, 192, 7, 106, 2, 2, 192, 193, 7,
	103, 2, 2, 193, 194, 7, 117, 2, 2, 194, 12, 3, 2, 2, 2, 195, 196, 7, 107,
	2, 2, 196, 197, 7, 112, 2, 2, 197, 14, 3, 2, 2, 2, 198, 199, 7, 112, 2,
	2, 199, 200, 7, 113, 2, 2, 200, 201, 7, 118, 2, 2, 201, 202, 7, 34, 2,
	2, 202, 203, 7, 107, 2, 2, 203, 204, 7, 112, 2, 2, 204, 16, 3, 2, 2, 2,
	205, 206, 7, 110, 2, 2, 206, 207, 7, 103, 2, 2, 207, 208, 7, 112, 2, 2,
	208, 18, 3, 2, 2, 2, 209, 210, 7, 99, 2, 2, 210, 211, 7, 110, 2, 2, 211,
	212, 7, 110, 2, 2, 212, 20, 3, 2, 2, 2, 213, 214, 7, 112, 2, 2, 214, 215,
	7, 113, 2, 2, 215, 216, 7, 112, 2, 2, 216, 217, 7, 103, 2, 2, 217, 22,
	3, 2, 2, 2, 218, 219, 7, 99, 2, 2, 219, 220, 7, 112, 2, 2, 220, 221, 7,
	123, 2, 2, 221, 24, 3, 2, 2, 2, 222, 223, 7, 113, 2, 2, 223, 224, 7, 112,
	2, 2, 224, 225, 7, 103, 2, 2, 225, 26, 3, 2, 2, 2, 226, 227, 7, 104, 2,
	2, 227, 228, 7, 107, 2, 2, 228, 229, 7, 110, 2, 2, 229, 230, 7, 118, 2,
	2, 230, 231, 7, 103, 2, 2, 231, 232, 7, 116, 2, 2, 232, 28, 3, 2, 2, 2,
	233, 234, 7, 111, 2, 2, 234, 235, 7, 99, 2, 2, 235, 236, 7, 114, 2, 2,
	236, 30, 3, 2, 2, 2, 237, 238, 7, 93, 2, 2, 238, 32, 3, 2, 2, 2, 239, 240,
	7, 95, 2, 2, 240, 34, 3, 2, 2, 2, 241, 242, 7, 42, 2, 2, 242, 36, 3, 2,
	2, 2, 243, 244, 7, 43, 2, 2, 244, 38, 3, 2, 2, 2, 245, 246, 7, 125, 2,
	2, 246, 40, 3, 2, 2, 2, 247, 248, 7, 127, 2, 2, 248, 42, 3, 2, 2, 2, 249,
	250, 7, 61, 2, 2, 250, 44, 3, 2, 2, 2, 251, 252, 7, 46, 2, 2, 252, 46,
	3, 2, 2, 2, 253, 254, 7, 63, 2, 2, 254, 48, 3, 2, 2, 2, 255, 256, 7, 65,
	2, 2, 256, 50, 3, 2, 2, 2, 257, 258, 7, 60, 2, 2, 258, 52, 3, 2, 2, 2,
	259, 260, 7, 48, 2, 2, 260, 54, 3, 2, 2, 2, 261, 262, 7, 45, 2, 2, 262,
	56, 3, 2, 2, 2, 263, 264, 7, 47, 2, 2, 264, 58, 3, 2, 2, 2, 265, 266, 7,
	35, 2, 2, 266, 60, 3, 2, 2, 2, 267, 268, 7, 112, 2, 2, 268, 269, 7, 113,
	2, 2, 269, 270, 7, 118, 2, 2, 270, 62, 3, 2, 2, 2, 271, 272, 7, 112, 2,
	2, 272, 273, 7, 107, 2, 2, 273, 274, 7, 110, 2, 2, 274, 64, 3, 2, 2, 2,
	275, 276, 7, 44, 2, 2, 276, 66, 3, 2, 2, 2, 277, 278, 7, 44, 2, 2, 278,
	279, 7, 44, 2, 2, 279, 68, 3, 2, 2, 2, 280, 281, 7, 49, 2, 2, 281, 70,
	3, 2, 2, 2, 282, 283, 7, 39, 2, 2, 283, 72, 3, 2, 2, 2, 284, 285, 7, 64,
	2, 2, 285, 286, 7, 64, 2, 2, 286, 74, 3, 2, 2, 2, 287, 288, 7, 62, 2, 2,
	288, 289, 7, 62, 2, 2, 289, 76, 3, 2, 2, 2, 290, 291, 7, 62, 2, 2, 291,
	78, 3, 2, 2, 2, 292, 293, 7, 64, 2, 2, 293, 80, 3, 2, 2, 2, 294, 295, 7,
	62, 2, 2, 295, 296, 7, 63, 2, 2, 296, 82, 3, 2, 2, 2, 297, 298, 7, 64,
	2, 2, 298, 299, 7, 63, 2, 2, 299, 84, 3, 2, 2, 2, 300, 301, 7, 63, 2, 2,
	301, 302, 7, 63, 2, 2, 302, 86, 3, 2, 2, 2, 303, 304, 7, 35, 2, 2, 304,
	305, 7, 63, 2, 2, 305, 88, 3, 2, 2, 2, 306, 307, 7, 37, 2, 2, 307, 90,
	3, 2, 2, 2, 308, 309, 7, 40, 2, 2, 309, 314, 7, 40, 2, 2, 310, 311, 7,
	99, 2, 2, 311, 312, 7, 112, 2, 2, 312, 314, 7, 102, 2, 2, 313, 308, 3,
	2, 2, 2, 313, 310, 3, 2, 2, 2, 314, 92, 3, 2, 2, 2, 315, 316, 7, 126, 2,
	2, 316, 320, 7, 126, 2, 2, 317, 318, 7, 113, 2, 2, 318, 320, 7, 116, 2,
	2, 319, 315, 3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 320, 94, 3, 2, 2, 2, 321,
	322, 7, 99, 2, 2, 322, 323, 7, 110, 2, 2, 323, 344, 7, 110, 2, 2, 324,
	325, 7, 112, 2, 2, 325, 326, 7, 113, 2, 2, 326, 327, 7, 112, 2, 2, 327,
	344, 7, 103, 2, 2, 328, 329, 7, 99, 2, 2, 329, 330, 7, 112, 2, 2, 330,
	344, 7, 123, 2, 2, 331, 332, 7, 113, 2, 2, 332, 333, 7, 112, 2, 2, 333,
	344, 7, 103, 2, 2, 334, 335, 7, 104, 2, 2, 335, 336, 7, 107, 2, 2, 336,
	337, 7, 110, 2, 2, 337, 338, 7, 118, 2, 2, 338, 339, 7, 103, 2, 2, 339,
	344, 7, 116, 2, 2, 340, 341, 7, 111, 2, 2, 341, 342, 7, 99, 2, 2, 342,
	344, 7, 114, 2, 2, 343, 321, 3, 2, 2, 2, 343, 324, 3, 2, 2, 2, 343, 328,
	3, 2, 2, 2, 343, 331, 3, 2, 2, 2, 343, 334, 3, 2, 2, 2, 343, 340, 3, 2,
	2, 2, 344, 96, 3, 2, 2, 2, 345, 346, 7, 117, 2, 2, 346, 347, 7, 118, 2,
	2, 347, 348, 7, 99, 2, 2, 348, 349, 7, 116, 2, 2, 349, 350, 7, 118, 2,
	2, 350, 351, 7, 117, 2, 2, 351, 352, 7, 89, 2, 2, 352, 353, 7, 107, 2,
	2, 353, 354, 7, 118, 2, 2, 354, 387, 7, 106, 2, 2, 355, 356, 7, 103, 2,
	2, 356, 357, 7, 112, 2, 2, 357, 358, 7, 102, 2, 2, 358, 359, 7, 117, 2,
	2, 359, 360, 7, 89, 2, 2, 360, 361, 7, 107, 2, 2, 361, 362, 7, 118, 2,
	2, 362, 387, 7, 106, 2, 2, 363, 364, 7, 101, 2, 2, 364, 365, 7, 113, 2,
	2, 365, 366, 7, 112, 2, 2, 366, 367, 7, 118, 2, 2, 367, 368, 7, 99, 2,
	2, 368, 369, 7, 107, 2, 2, 369, 370, 7, 112, 2, 2, 370, 387, 7, 117, 2,
	2, 371, 372, 7, 111, 2, 2, 372, 373, 7, 99, 2, 2, 373, 374, 7, 118, 2,
	2, 374, 375, 7, 101, 2, 2, 375, 376, 7, 106, 2, 2, 376, 377, 7, 103, 2,
	2, 377, 387, 7, 117, 2, 2, 378, 379, 7, 107, 2, 2, 379, 387, 7, 112, 2,
	2, 380, 381, 7, 112, 2, 2, 381, 382, 7, 113, 2, 2, 382, 383, 7, 118, 2,
	2, 383, 384, 7, 34, 2, 2, 384, 385, 7, 107, 2, 2, 385, 387, 7, 112, 2,
	2, 386, 345, 3, 2, 2, 2, 386, 355, 3, 2, 2, 2, 386, 363, 3, 2, 2, 2, 386,
	371, 3, 2, 2, 2, 386, 378, 3, 2, 2, 2, 386, 380, 3, 2, 2, 2, 387, 98, 3,
	2, 2, 2, 388, 389, 7, 118, 2, 2, 389, 390, 7, 116, 2, 2, 390, 391, 7, 119,
	2, 2, 391, 398, 7, 103, 2, 2, 392, 393, 7, 104, 2, 2, 393, 394, 7, 99,
	2, 2, 394, 395, 7, 110, 2, 2, 395, 396, 7, 117, 2, 2, 396, 398, 7, 103,
	2, 2, 397, 388, 3, 2, 2, 2, 397, 392, 3, 2, 2, 2, 398, 100, 3, 2, 2, 2,
	399, 408, 7, 50, 2, 2, 400, 404, 9, 2, 2, 2, 401, 403, 9, 3, 2, 2, 402,
	401, 3, 2, 2, 2, 403, 406, 3, 2, 2, 2, 404, 402, 3, 2, 2, 2, 404, 405,
	3, 2, 2, 2, 405, 408, 3, 2, 2, 2, 406, 404, 3, 2, 2, 2, 407, 399, 3, 2,
	2, 2, 407, 400, 3, 2, 2, 2, 408, 102, 3, 2, 2, 2, 409, 410, 5, 151, 76,
	2, 410, 412, 7, 48, 2, 2, 411, 413, 5, 149, 75, 2, 412, 411, 3, 2, 2, 2,
	413, 414, 3, 2, 2, 2, 414, 412, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415,
	423, 3, 2, 2, 2, 416, 418, 7, 48, 2, 2, 417, 419, 5, 149, 75, 2, 418, 417,
	3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 418, 3, 2, 2, 2, 420, 421, 3, 2,
	2, 2, 421, 423, 3, 2, 2, 2, 422, 409, 3, 2, 2, 2, 422, 416, 3, 2, 2, 2,
	423, 104, 3, 2, 2, 2, 424, 425, 7, 50, 2, 2, 425, 427, 9, 4, 2, 2, 426,
	428, 5, 153, 77, 2, 427, 426, 3, 2, 2, 2, 428, 429, 3, 2, 2, 2, 429, 427,
	3, 2, 2, 2, 429, 430, 3, 2, 2, 2, 430, 106, 3, 2, 2, 2, 431, 435, 5, 143,
	72, 2, 432, 434, 5, 145, 73, 2, 433, 432, 3, 2, 2, 2, 434, 437, 3, 2, 2,
	2, 435, 433, 3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436, 108, 3, 2, 2, 2, 437,
	435, 3, 2, 2, 2, 438, 442, 7, 36, 2, 2, 439, 441, 5, 121, 61, 2, 440, 439,
	3, 2, 2, 2, 441, 444, 3, 2, 2, 2, 442, 440, 3, 2, 2, 2, 442, 443, 3, 2,
	2, 2, 443, 445, 3, 2, 2, 2, 444, 442, 3, 2, 2, 2, 445, 455, 7, 36, 2, 2,
	446, 450, 7, 41, 2, 2, 447, 449, 5, 123, 62, 2, 448, 447, 3, 2, 2, 2, 449,
	452, 3, 2, 2, 2, 450, 448, 3, 2, 2, 2, 450, 451, 3, 2, 2, 2, 451, 453,
	3, 2, 2, 2, 452, 450, 3, 2, 2, 2, 453, 455, 7, 41, 2, 2, 454, 438, 3, 2,
	2, 2, 454, 446, 3, 2, 2, 2, 455, 110, 3, 2, 2, 2, 456, 458, 9, 5, 2, 2,
	457, 456, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 457, 3, 2, 2, 2, 459,
	460, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461, 462, 8, 56, 2, 2, 462, 112,
	3, 2, 2, 2, 463, 464, 7, 49, 2, 2, 464, 465, 7, 44, 2, 2, 465, 469, 3,
	2, 2, 2, 466, 468, 11, 2, 2, 2, 467, 466, 3, 2, 2, 2, 468, 471, 3, 2, 2,
	2, 469, 470, 3, 2, 2, 2, 469, 467, 3, 2, 2, 2, 470, 472, 3, 2, 2, 2, 471,
	469, 3, 2, 2, 2, 472, 473, 7, 44, 2, 2, 473, 474, 7, 49, 2, 2, 474, 475,
	3, 2, 2, 2, 475, 476, 8, 57, 2, 2, 476, 114, 3, 2, 2, 2, 477, 478, 7, 49,
	2, 2, 478, 479, 7, 49, 2, 2, 479, 483, 3, 2, 2, 2, 480, 482, 10, 6, 2,
	2, 481, 480, 3, 2, 2, 2, 482, 485, 3, 2, 2, 2, 483, 481, 3, 2, 2, 2, 483,
	484, 3, 2, 2, 2, 484, 486, 3, 2, 2, 2, 485, 483, 3, 2, 2, 2, 486, 487,
	8, 58, 2, 2, 487, 116, 3, 2, 2, 2, 488, 489, 9, 6, 2, 2, 489, 490, 3, 2,
	2, 2, 490, 491, 8, 59, 2, 2, 491, 118, 3, 2, 2, 2, 492, 493, 11, 2, 2,
	2, 493, 120, 3, 2, 2, 2, 494, 499, 10, 7, 2, 2, 495, 496, 7, 94, 2, 2,
	496, 499, 5, 125, 63, 2, 497, 499, 5, 139, 70, 2, 498, 494, 3, 2, 2, 2,
	498, 495, 3, 2, 2, 2, 498, 497, 3, 2, 2, 2, 499, 122, 3, 2, 2, 2, 500,
	505, 10, 8, 2, 2, 501, 502, 7, 94, 2, 2, 502, 505, 5, 125, 63, 2, 503,
	505, 5, 139, 70, 2, 504, 500, 3, 2, 2, 2, 504, 501, 3, 2, 2, 2, 504, 503,
	3, 2, 2, 2, 505, 124, 3, 2, 2, 2, 506, 511, 5, 127, 64, 2, 507, 511, 7,
	50, 2, 2, 508, 511, 5, 129, 65, 2, 509, 511, 5, 131, 66, 2, 510, 506, 3,
	2, 2, 2, 510, 507, 3, 2, 2, 2, 510, 508, 3, 2, 2, 2, 510, 509, 3, 2, 2,
	2, 511, 126, 3, 2, 2, 2, 512, 515, 5, 133, 67, 2, 513, 515, 5, 135, 68,
	2, 514, 512, 3, 2, 2, 2, 514, 513, 3, 2, 2, 2, 515, 128, 3, 2, 2, 2, 516,
	517, 7, 122, 2, 2, 517, 518, 5, 153, 77, 2, 518, 519, 5, 153, 77, 2, 519,
	130, 3, 2, 2, 2, 520, 521, 7, 119, 2, 2, 521, 522, 5, 153, 77, 2, 522,
	523, 5, 153, 77, 2, 523, 524, 5, 153, 77, 2, 524, 525, 5, 153, 77, 2, 525,
	132, 3, 2, 2, 2, 526, 527, 9, 9, 2, 2, 527, 134, 3, 2, 2, 2, 528, 529,
	10, 10, 2, 2, 529, 136, 3, 2, 2, 2, 530, 534, 5, 133, 67, 2, 531, 534,
	5, 149, 75, 2, 532, 534, 9, 11, 2, 2, 533, 530, 3, 2, 2, 2, 533, 531, 3,
	2, 2, 2, 533, 532, 3, 2, 2, 2, 534, 138, 3, 2, 2, 2, 535, 536, 7, 94, 2,
	2, 536, 537, 5, 141, 71, 2, 537, 140, 3, 2, 2, 2, 538, 539, 7, 15, 2, 2,
	539, 542, 7, 12, 2, 2, 540, 542, 5, 117, 59, 2, 541, 538, 3, 2, 2, 2, 541,
	540, 3, 2, 2, 2, 542, 142, 3, 2, 2, 2, 543, 546, 5, 147, 74, 2, 544, 546,
	9, 12, 2, 2, 545, 543, 3, 2, 2, 2, 545, 544, 3, 2, 2, 2, 546, 144, 3, 2,
	2, 2, 547, 551, 5, 147, 74, 2, 548, 551, 5, 149, 75, 2, 549, 551, 9, 13,
	2, 2, 550, 547, 3, 2, 2, 2, 550, 548, 3, 2, 2, 2, 550, 549, 3, 2, 2, 2,
	551, 146, 3, 2, 2, 2, 552, 553, 9, 14, 2, 2, 553, 148, 3, 2, 2, 2, 554,
	555, 4, 50, 59, 2, 555, 150, 3, 2, 2, 2, 556, 565, 7, 50, 2, 2, 557, 561,
	9, 2, 2, 2, 558, 560, 5, 149, 75, 2, 559, 558, 3, 2, 2, 2, 560, 563, 3,
	2, 2, 2, 561, 559, 3, 2, 2, 2, 561, 562, 3, 2, 2, 2, 562, 565, 3, 2, 2,
	2, 563, 561, 3, 2, 2, 2, 564, 556, 3, 2, 2, 2, 564, 557, 3, 2, 2, 2, 565,
	152, 3, 2, 2, 2, 566, 567, 9, 15, 2, 2, 567, 154, 3, 2, 2, 2, 31, 2, 313,
	319, 343, 386, 397, 404, 407, 414, 420, 422, 429, 435, 442, 450, 454, 459,
	469, 483, 498, 504, 510, 514, 533, 541, 545, 550, 561, 564, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'..'", "'startsWith'", "'endsWith'", "'contains'", "'matches'", "'in'",
	"'not in'", "'len'", "'all'", "'none'", "'any'", "'one'", "'filter'", "'map'",
	"'['", "']'", "'('", "')'", "'{'", "'}'", "';'", "','", "'='", "'?'", "':'",
	"'.'", "'+'", "'-'", "'!'", "'not'", "'nil'", "'*'", "'**'", "'/'", "'%'",
	"'>>'", "'<<'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'#'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "OpenBracket",
	"CloseBracket", "OpenParen", "CloseParen", "OpenBrace", "CloseBrace", "SemiColon",
	"Comma", "Assign", "QuestionMark", "Colon", "Dot", "Plus", "Minus", "Negate",
	"Not", "Nil", "Multiply", "Exponent", "Divide", "Modulus", "RightShiftArithmetic",
	"LeftShiftArithmetic", "LessThan", "MoreThan", "LessThanEquals", "GreaterThanEquals",
	"Equals", "NotEquals", "Pointer", "And", "Or", "Builtins", "Ops", "BooleanLiteral",
	"IntegerLiteral", "FloatLiteral", "HexIntegerLiteral", "Identifier", "StringLiteral",
	"WhiteSpaces", "MultiLineComment", "SingleLineComment", "LineTerminator",
	"UnexpectedCharacter",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "OpenBracket", "CloseBracket",
	"OpenParen", "CloseParen", "OpenBrace", "CloseBrace", "SemiColon", "Comma",
	"Assign", "QuestionMark", "Colon", "Dot", "Plus", "Minus", "Negate", "Not",
	"Nil", "Multiply", "Exponent", "Divide", "Modulus", "RightShiftArithmetic",
	"LeftShiftArithmetic", "LessThan", "MoreThan", "LessThanEquals", "GreaterThanEquals",
	"Equals", "NotEquals", "Pointer", "And", "Or", "Builtins", "Ops", "BooleanLiteral",
	"IntegerLiteral", "FloatLiteral", "HexIntegerLiteral", "Identifier", "StringLiteral",
	"WhiteSpaces", "MultiLineComment", "SingleLineComment", "LineTerminator",
	"UnexpectedCharacter", "DoubleStringCharacter", "SingleStringCharacter",
	"EscapeSequence", "CharacterEscapeSequence", "HexEscapeSequence", "UnicodeEscapeSequence",
	"SingleEscapeCharacter", "NonEscapeCharacter", "EscapeCharacter", "LineContinuation",
	"LineTerminatorSequence", "IdentifierStart", "IdentifierPart", "Letter",
	"Digit", "DecimalLiteral", "HexDigit",
}

type ExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewExprLexer(input antlr.CharStream) *ExprLexer {

	l := new(ExprLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Expr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExprLexer tokens.
const (
	ExprLexerT__0                 = 1
	ExprLexerT__1                 = 2
	ExprLexerT__2                 = 3
	ExprLexerT__3                 = 4
	ExprLexerT__4                 = 5
	ExprLexerT__5                 = 6
	ExprLexerT__6                 = 7
	ExprLexerT__7                 = 8
	ExprLexerT__8                 = 9
	ExprLexerT__9                 = 10
	ExprLexerT__10                = 11
	ExprLexerT__11                = 12
	ExprLexerT__12                = 13
	ExprLexerT__13                = 14
	ExprLexerOpenBracket          = 15
	ExprLexerCloseBracket         = 16
	ExprLexerOpenParen            = 17
	ExprLexerCloseParen           = 18
	ExprLexerOpenBrace            = 19
	ExprLexerCloseBrace           = 20
	ExprLexerSemiColon            = 21
	ExprLexerComma                = 22
	ExprLexerAssign               = 23
	ExprLexerQuestionMark         = 24
	ExprLexerColon                = 25
	ExprLexerDot                  = 26
	ExprLexerPlus                 = 27
	ExprLexerMinus                = 28
	ExprLexerNegate               = 29
	ExprLexerNot                  = 30
	ExprLexerNil                  = 31
	ExprLexerMultiply             = 32
	ExprLexerExponent             = 33
	ExprLexerDivide               = 34
	ExprLexerModulus              = 35
	ExprLexerRightShiftArithmetic = 36
	ExprLexerLeftShiftArithmetic  = 37
	ExprLexerLessThan             = 38
	ExprLexerMoreThan             = 39
	ExprLexerLessThanEquals       = 40
	ExprLexerGreaterThanEquals    = 41
	ExprLexerEquals               = 42
	ExprLexerNotEquals            = 43
	ExprLexerPointer              = 44
	ExprLexerAnd                  = 45
	ExprLexerOr                   = 46
	ExprLexerBuiltins             = 47
	ExprLexerOps                  = 48
	ExprLexerBooleanLiteral       = 49
	ExprLexerIntegerLiteral       = 50
	ExprLexerFloatLiteral         = 51
	ExprLexerHexIntegerLiteral    = 52
	ExprLexerIdentifier           = 53
	ExprLexerStringLiteral        = 54
	ExprLexerWhiteSpaces          = 55
	ExprLexerMultiLineComment     = 56
	ExprLexerSingleLineComment    = 57
	ExprLexerLineTerminator       = 58
	ExprLexerUnexpectedCharacter  = 59
)