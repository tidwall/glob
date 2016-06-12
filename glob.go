//
// The code in this file is derived from the Redis project.
// Specifically from:
//    FILE:   https://github.com/antirez/redis/blob/unstable/src/util.c
//    COMMIT: 6dead2c
//
// Original author: Salvatore Sanfilippo
// Modified by: Josh Baker
//
// The verbatim license is below.
//

/*
 * Copyright (c) 2009-2012, Salvatore Sanfilippo <antirez at gmail dot com>
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *   * Redistributions of source code must retain the above copyright notice,
 *     this list of conditions and the following disclaimer.
 *   * Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in the
 *     documentation and/or other materials provided with the distribution.
 *   * Neither the name of Redis nor the names of its contributors may be used
 *     to endorse or promote products derived from this software without
 *     specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

package glob

/* Glob-style pattern matching. */
func Match(pattern, str string, nocase bool) bool {
	for len(pattern) > 0 {
		switch pattern[0] {
		case '*':
			for len(pattern) > 1 && pattern[1] == '*' {
				pattern = pattern[1:]
			}
			if len(pattern) == 1 {
				return true /* match */
			}
			for len(str) > 0 {
				if Match(pattern[1:], str, nocase) {
					return true /* match */
				}
				str = str[1:]
			}
			return false /* no match */
		case '?':
			if len(str) == 0 {
				return false /* no match */
			}
			str = str[1:]
		case '[':
			var not, match bool
			pattern = pattern[1:]
			not = pattern[0] == '^'
			if not {
				pattern = pattern[1:]
			}
			for {
				if pattern[0] == '\\' {
					pattern = pattern[1:]
					if pattern[0] == str[0] {
						match = true
					}
				} else if pattern[0] == ']' {
					break
				} else if len(pattern) == 0 {
					pattern = pattern[1:]
					break
				} else if pattern[1] == '-' && len(pattern) >= 3 {
					var start = pattern[0]
					var end = pattern[2]
					var c = str[0]
					if start > end {
						start, end = end, start
					}
					if nocase {
						start = tolower(start)
						end = tolower(end)
						c = tolower(c)
					}
					pattern = pattern[2:]
					if c >= start && c <= end {
						match = true
					}
				} else {
					if !nocase {
						if pattern[0] == str[0] {
							match = true
						}
					} else {
						if tolower(pattern[0]) == tolower(str[0]) {
							match = true
						}
					}
				}
				pattern = pattern[1:]
			}
			if not {
				match = !match
			}
			if !match {
				return false /* no match */
			}
			str = str[1:]
		case '\\':
			if len(pattern) >= 2 {
				pattern = pattern[1:]
			}
			fallthrough
			/* fall through */
		default:
			if !nocase {
				if pattern[0] != str[0] {
					return false /* no match */
				}
			} else {
				if tolower(pattern[0]) != tolower(str[0]) {
					return false /* no match */
				}
			}
			str = str[1:]
			break
		}
		pattern = pattern[1:]
		if len(str) == 0 {
			for len(pattern) > 0 && pattern[0] == '*' {
				pattern = pattern[1:]
			}
			break
		}
	}
	if len(pattern) == 0 && len(str) == 0 {
		return true
	}
	return false
}

func tolower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
