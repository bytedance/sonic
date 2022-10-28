/*
 * Copyright 2021 ByteDance Inc.
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

package issue_test

import (
    `bytes`
    `compress/gzip`
    `encoding/json`
    `io/ioutil`
    `reflect`
    `testing`
    `time`

    . `github.com/bytedance/sonic`
    `github.com/bytedance/sonic/option`
    `github.com/stretchr/testify/require`
)

var jsonData = func() string {
    // Read and decompress the test data.
    b, err := ioutil.ReadFile("../testdata/synthea_fhir.json.gz")
    if err != nil {
        panic(err)
    }
    zr, err := gzip.NewReader(bytes.NewReader(b))
    if err != nil {
        panic(err)
    }
    data, err := ioutil.ReadAll(zr)
    if err != nil {
        panic(err)
    }
    return string(data)
}()

type (
    syntheaRoot struct {
        Entry []struct {
            FullURL string `json:"fullUrl"`
            Request *struct {
                Method string `json:"method"`
                URL    string `json:"url"`
            } `json:"request"`
            Resource *struct {
                AbatementDateTime time.Time   `json:"abatementDateTime"`
                AchievementStatus syntheaCode `json:"achievementStatus"`
                Active            bool        `json:"active"`
                Activity          []struct {
                    Detail *struct {
                        Code     syntheaCode      `json:"code"`
                        Location syntheaReference `json:"location"`
                        Status   string           `json:"status"`
                    } `json:"detail"`
                } `json:"activity"`
                Address        []syntheaAddress   `json:"address"`
                Addresses      []syntheaReference `json:"addresses"`
                AuthoredOn     time.Time          `json:"authoredOn"`
                BillablePeriod syntheaRange       `json:"billablePeriod"`
                BirthDate      string             `json:"birthDate"`
                CareTeam       []struct {
                    Provider  syntheaReference `json:"provider"`
                    Reference string           `json:"reference"`
                    Role      syntheaCode      `json:"role"`
                    Sequence  int64            `json:"sequence"`
                } `json:"careTeam"`
                Category       []syntheaCode    `json:"category"`
                Claim          syntheaReference `json:"claim"`
                Class          syntheaCoding    `json:"class"`
                ClinicalStatus syntheaCode      `json:"clinicalStatus"`
                Code           syntheaCode      `json:"code"`
                Communication  []struct {
                    Language syntheaCode `json:"language"`
                } `json:"communication"`
                Component []struct {
                    Code          syntheaCode   `json:"code"`
                    ValueQuantity syntheaCoding `json:"valueQuantity"`
                } `json:"component"`
                Contained []struct {
                    Beneficiary  syntheaReference   `json:"beneficiary"`
                    ID           string             `json:"id"`
                    Intent       string             `json:"intent"`
                    Payor        []syntheaReference `json:"payor"`
                    Performer    []syntheaReference `json:"performer"`
                    Requester    syntheaReference   `json:"requester"`
                    ResourceType string             `json:"resourceType"`
                    Status       string             `json:"status"`
                    Subject      syntheaReference   `json:"subject"`
                    Type         syntheaCode        `json:"type"`
                } `json:"contained"`
                Created          time.Time   `json:"created"`
                DeceasedDateTime time.Time   `json:"deceasedDateTime"`
                Description      syntheaCode `json:"description"`
                Diagnosis        []struct {
                    DiagnosisReference syntheaReference `json:"diagnosisReference"`
                    Sequence           int64            `json:"sequence"`
                    Type               []syntheaCode    `json:"type"`
                } `json:"diagnosis"`
                DosageInstruction []struct {
                    AsNeededBoolean bool `json:"asNeededBoolean"`
                    DoseAndRate     []struct {
                        DoseQuantity *struct {
                            Value float64 `json:"value"`
                        } `json:"doseQuantity"`
                        Type syntheaCode `json:"type"`
                    } `json:"doseAndRate"`
                    Sequence int64 `json:"sequence"`
                    Timing   *struct {
                        Repeat *struct {
                            Frequency  int64   `json:"frequency"`
                            Period     float64 `json:"period"`
                            PeriodUnit string  `json:"periodUnit"`
                        } `json:"repeat"`
                    } `json:"timing"`
                } `json:"dosageInstruction"`
                EffectiveDateTime time.Time          `json:"effectiveDateTime"`
                Encounter         syntheaReference   `json:"encounter"`
                Extension         []syntheaExtension `json:"extension"`
                Gender            string             `json:"gender"`
                Goal              []syntheaReference `json:"goal"`
                ID                string             `json:"id"`
                Identifier        []struct {
                    System string      `json:"system"`
                    Type   syntheaCode `json:"type"`
                    Use    string      `json:"use"`
                    Value  string      `json:"value"`
                } `json:"identifier"`
                Insurance []struct {
                    Coverage syntheaReference `json:"coverage"`
                    Focal    bool             `json:"focal"`
                    Sequence int64            `json:"sequence"`
                } `json:"insurance"`
                Insurer syntheaReference `json:"insurer"`
                Intent  string           `json:"intent"`
                Issued  time.Time        `json:"issued"`
                Item    []struct {
                    Adjudication []struct {
                        Amount   syntheaCurrency `json:"amount"`
                        Category syntheaCode     `json:"category"`
                    } `json:"adjudication"`
                    Category                syntheaCode        `json:"category"`
                    DiagnosisSequence       []int64            `json:"diagnosisSequence"`
                    Encounter               []syntheaReference `json:"encounter"`
                    InformationSequence     []int64            `json:"informationSequence"`
                    LocationCodeableConcept syntheaCode        `json:"locationCodeableConcept"`
                    Net                     syntheaCurrency    `json:"net"`
                    ProcedureSequence       []int64            `json:"procedureSequence"`
                    ProductOrService        syntheaCode        `json:"productOrService"`
                    Sequence                int64              `json:"sequence"`
                    ServicedPeriod          syntheaRange       `json:"servicedPeriod"`
                } `json:"item"`
                LifecycleStatus           string             `json:"lifecycleStatus"`
                ManagingOrganization      []syntheaReference `json:"managingOrganization"`
                MaritalStatus             syntheaCode        `json:"maritalStatus"`
                MedicationCodeableConcept syntheaCode        `json:"medicationCodeableConcept"`
                MultipleBirthBoolean      bool               `json:"multipleBirthBoolean"`
                Name                      json.RawMessage           `json:"name"`
                NumberOfInstances         int64              `json:"numberOfInstances"`
                NumberOfSeries            int64              `json:"numberOfSeries"`
                OccurrenceDateTime        time.Time          `json:"occurrenceDateTime"`
                OnsetDateTime             time.Time          `json:"onsetDateTime"`
                Outcome                   string             `json:"outcome"`
                Participant               []struct {
                    Individual syntheaReference `json:"individual"`
                    Member     syntheaReference `json:"member"`
                    Role       []syntheaCode    `json:"role"`
                } `json:"participant"`
                Patient syntheaReference `json:"patient"`
                Payment *struct {
                    Amount syntheaCurrency `json:"amount"`
                } `json:"payment"`
                PerformedPeriod syntheaRange     `json:"performedPeriod"`
                Period          syntheaRange     `json:"period"`
                Prescription    syntheaReference `json:"prescription"`
                PrimarySource   bool             `json:"primarySource"`
                Priority        syntheaCode      `json:"priority"`
                Procedure       []struct {
                    ProcedureReference syntheaReference `json:"procedureReference"`
                    Sequence           int64            `json:"sequence"`
                } `json:"procedure"`
                Provider        syntheaReference   `json:"provider"`
                ReasonCode      []syntheaCode      `json:"reasonCode"`
                ReasonReference []syntheaReference `json:"reasonReference"`
                RecordedDate    time.Time          `json:"recordedDate"`
                Referral        syntheaReference   `json:"referral"`
                Requester       syntheaReference   `json:"requester"`
                ResourceType    string             `json:"resourceType"`
                Result          []syntheaReference `json:"result"`
                Series          []struct {
                    BodySite syntheaCoding `json:"bodySite"`
                    Instance []struct {
                        Number   int64         `json:"number"`
                        SopClass syntheaCoding `json:"sopClass"`
                        Title    string        `json:"title"`
                        UID      string        `json:"uid"`
                    } `json:"instance"`
                    Modality          syntheaCoding `json:"modality"`
                    Number            int64         `json:"number"`
                    NumberOfInstances int64         `json:"numberOfInstances"`
                    Started           string        `json:"started"`
                    UID               string        `json:"uid"`
                } `json:"series"`
                ServiceProvider syntheaReference `json:"serviceProvider"`
                Started         time.Time        `json:"started"`
                Status          string           `json:"status"`
                Subject         syntheaReference `json:"subject"`
                SupportingInfo  []struct {
                    Category       syntheaCode      `json:"category"`
                    Sequence       int64            `json:"sequence"`
                    ValueReference syntheaReference `json:"valueReference"`
                } `json:"supportingInfo"`
                Telecom              []map[string]string `json:"telecom"`
                Text                 map[string]string   `json:"text"`
                Total                json.RawMessage            `json:"total"`
                Type                 json.RawMessage            `json:"type"`
                Use                  string              `json:"use"`
                VaccineCode          syntheaCode         `json:"vaccineCode"`
                ValueCodeableConcept syntheaCode         `json:"valueCodeableConcept"`
                ValueQuantity        syntheaCoding       `json:"valueQuantity"`
                VerificationStatus   syntheaCode         `json:"verificationStatus"`
            } `json:"resource"`
        } `json:"entry"`
        ResourceType string `json:"resourceType"`
        Type         string `json:"type"`
    }
    syntheaCode struct {
        Coding []syntheaCoding `json:"coding"`
        Text   string          `json:"text"`
    }
    syntheaCoding struct {
        Code    string  `json:"code"`
        Display string  `json:"display"`
        System  string  `json:"system"`
        Unit    string  `json:"unit"`
        Value   float64 `json:"value"`
    }
    syntheaReference struct {
        Display   string `json:"display"`
        Reference string `json:"reference"`
    }
    syntheaAddress struct {
        City       string             `json:"city"`
        Country    string             `json:"country"`
        Extension  []syntheaExtension `json:"extension"`
        Line       []string           `json:"line"`
        PostalCode string             `json:"postalCode"`
        State      string             `json:"state"`
    }
    syntheaExtension struct {
        URL          string             `json:"url"`
        ValueAddress syntheaAddress     `json:"valueAddress"`
        ValueCode    string             `json:"valueCode"`
        ValueDecimal float64            `json:"valueDecimal"`
        ValueString  string             `json:"valueString"`
        Extension    []syntheaExtension `json:"extension"`
    }
    syntheaRange struct {
        End   time.Time `json:"end"`
        Start time.Time `json:"start"`
    }
    syntheaCurrency struct {
        Currency string  `json:"currency"`
        Value    float64 `json:"value"`
    }
)
 

func TestPretouchSynteaRoot(t *testing.T) {
    m := new(syntheaRoot)
    s := time.Now()
    println("start decoder pretouch:", s.UnixNano())
    require.Nil(t, Pretouch(reflect.TypeOf(m), option.WithCompileMaxInlineDepth(2), option.WithCompileRecursiveDepth(20)))
    e := time.Now()
    println("end decoder pretouch:", e.UnixNano())
    println("elapsed:", e.Sub(s).Milliseconds(), "ms")
    
    s = time.Now()
    println("start decode:", s.UnixNano())
    require.Nil(t, UnmarshalString(jsonData, m))
    e = time.Now()
    println("end decode:", e.UnixNano())
    d1 := e.Sub(s).Nanoseconds()
    println("elapsed:", d1, "ns")

    s = time.Now()
    println("start decode:", s.UnixNano())
    require.Nil(t, UnmarshalString(jsonData, m))
    e = time.Now()
    println("end decode:", e.UnixNano())
    d2 := e.Sub(s).Nanoseconds()
    println("elapsed:", d2, "ns")
    if d1 > d2 * 20 {
        t.Fatal("decoder pretouch not finish yet")
    }

    s = time.Now()
    println("start decode:", s.UnixNano())
    require.Nil(t, UnmarshalString(jsonData, m))
    e = time.Now()
    println("end decode:", e.UnixNano())
    d5 := e.Sub(s).Nanoseconds()
    println("elapsed:", d5, "ns")
    if d2 > d5 * 10 {
        t.Fatal("decoder pretouch not finish yet")
    }

    s = time.Now()
    println("start encode 1:", s.UnixNano())
    _, err := MarshalString(*m)
    require.Nil(t, err)
    e = time.Now()
    println("end encode 1:", e.UnixNano())
    d3 := e.Sub(s).Nanoseconds()
    println("elapsed:", d3, "ns")
    
    s = time.Now()
    println("start encode 2:", s.UnixNano())
    _, err = MarshalString(m)
    require.Nil(t, err)
    e = time.Now()
    println("end encode 2:", e.UnixNano())
    d4 := e.Sub(s).Nanoseconds()
    println("elapsed:", d4, "ns")
    // if d3 > d4 * 10 {
    //     t.Fatal("encoder pretouch not finish yet")
    // }

    s = time.Now()
    println("start encode 3:", s.UnixNano())
    _, err = MarshalString(m)
    require.Nil(t, err)
    e = time.Now()
    println("end encode 3:", e.UnixNano())
    d6 := e.Sub(s).Nanoseconds()
    println("elapsed:", d6, "ns")
    if d4 > d6 * 10 {
        t.Fatal("encoder pretouch not finish yet")
    }
}

func BenchmarkDecodeSynteaRoot(b *testing.B) {
    m := new(syntheaRoot)
    require.Nil(b, Pretouch(reflect.TypeOf(m), option.WithCompileRecursiveDepth(10)))

    b.SetBytes(int64(len(jsonData)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = UnmarshalString(jsonData, m)
    }
}

func BenchmarkEncodeSynteaRoot(b *testing.B) {
    m := new(syntheaRoot)
    require.Nil(b, Pretouch(reflect.TypeOf(m), option.WithCompileRecursiveDepth(10)))
    require.Nil(b, UnmarshalString(jsonData, m))

    b.SetBytes(int64(len(jsonData)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = MarshalString(m)
    }
}