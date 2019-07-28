// Copyright © 2019 Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gen

import (
	"fmt"
	"math/rand"
	"time"
)

func Bmi(donorIdx int, date time.Time, value float64) Object {
	return Object{
		"resourceType":      "Observation",
		"id":                fmt.Sprintf("%d-bmi", donorIdx),
		"meta":              meta("https://fhir.bbmri.de/StructureDefinition/BbmriBmi"),
		"status":            "final",
		"category":          Array{vitalSigns},
		"subject":           reference("Patient", donorIdx),
		"code":              codeableConcept(coding("http://loinc.org", "39156-5")),
		"effectiveDateTime": date.Format("2006-01-02"),
		"valueQuantity":     quantity(value, "kg/m2"),
	}
}

func RandBmiValue(r *rand.Rand) float64 {
	return r.NormFloat64()*4.89 + 27.3
}
