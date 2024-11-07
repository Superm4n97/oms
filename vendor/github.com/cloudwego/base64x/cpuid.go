/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package base64x

import (
    `fmt`
    `os`

    `github.com/klauspost/cpuid/v2`
)

func hasAVX2() bool {
    switch v := os.Getenv("B64X_MODE"); v {
        case ""       : fallthrough
        case "auto"   : return cpuid.CPU.Has(cpuid.AVX2)
        case "noavx2" : return false
        default       : panic(fmt.Sprintf("invalid mode: '%s', should be one of 'auto', 'noavx2'", v))
    }
}
