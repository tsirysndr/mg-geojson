package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/rs/xid"
)

const (
	country      = "data/mdg_admbnda_adm0_BNGRC_OCHA_20181031.json"
	regions      = "data/mdg_admbnda_adm1_BNGRC_OCHA_20181031.json"
	districts    = "data/mdg_admbnda_adm2_BNGRC_OCHA_20181031.json"
	communes     = "data/mdg_admbnda_adm3_BNGRC_OCHA_20181031.json"
	fokontany    = "data/mdg_admbnda_adm4_BNGRC_OCHA_20181031.json"
	countryOut   = "assets/country.json"
	regionsOut   = "assets/regions.json"
	districtsOut = "assets/districts.json"
	communesOut  = "assets/communes.json"
	fokontanyOut = "assets/fokontany.json"
)

type Properties struct {
	Adm0Pcode string `json:"ADM0_PCODE"`
	Adm0En    string `json:"ADM0_EN"`
	Adm1Pcode string `json:"ADM1_PCODE"`
	Adm1En    string `json:"ADM1_EN"`
	Adm1Type  string `json:"ADM1_TYPE"`
	Adm2Pcode string `json:"ADM2_PCODE"`
	Adm2En    string `json:"ADM2_EN"`
	Adm2Type  string `json:"ADM2_TYPE"`
	Adm3Pcode string `json:"ADM3_PCODE"`
	Adm3En    string `json:"ADM3_EN"`
	Adm3Type  string `json:"ADM3_TYPE"`
	Adm4Pcode string `json:"ADM4_PCODE"`
	Adm4En    string `json:"ADM4_EN"`
	Adm4Type  string `json:"ADM4_TYPE"`
	ProvCode  int    `json:"PROV_CODE_"`
	OldProvin string `json:"OLD_PROVIN"`
}

type Geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

type Feature struct {
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

type GeoJSON struct {
	Features []Feature `json:"features"`
}

type Country struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Code     string   `json:"code"`
	Geometry Geometry `json:"geometry"`
}

type Region struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Code     string   `json:"code"`
	Province string   `json:"province"`
	Geometry Geometry `json:"geometry"`
}

type District struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Code     string   `json:"code"`
	Province string   `json:"province"`
	Geometry Geometry `json:"geometry"`
}

type Commune struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Code     string   `json:"code"`
	Province string   `json:"province"`
	Geometry Geometry `json:"geometry"`
}

type Fokontany struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Code     string   `json:"code"`
	Province string   `json:"province"`
	Geometry Geometry `json:"geometry"`
}

type Result struct {
	filename string
	GeoJSON  *GeoJSON
}

func Start() {
	countryCh := make(chan *Result)
	regionsCh := make(chan *Result)
	districtsCh := make(chan *Result)
	communesCh := make(chan *Result)
	fokontanyCh := make(chan *Result)

	go ParseGeoJSON(country, countryCh)
	go ParseGeoJSON(regions, regionsCh)
	go ParseGeoJSON(districts, districtsCh)
	go ParseGeoJSON(communes, communesCh)
	go ParseGeoJSON(fokontany, fokontanyCh)

	countryRes := <-countryCh
	regionsRes := <-regionsCh
	districtsRes := <-districtsCh
	communesRes := <-communesCh
	fokontanyRes := <-fokontanyCh

	log.Printf("%s file parsed ...", countryRes.filename)
	log.Printf("%s file parsed ...", regionsRes.filename)
	log.Printf("%s file parsed ...", districtsRes.filename)
	log.Printf("%s file parsed ...", communesRes.filename)
	log.Printf("%s file parsed ...", fokontanyRes.filename)

	ConvertToCountry(countryRes.GeoJSON)
	ConvertToRegions(regionsRes.GeoJSON)
	ConvertToDistricts(districtsRes.GeoJSON)
	ConvertToCommunes(communesRes.GeoJSON)
	ConvertToFokontany(fokontanyRes.GeoJSON)
}

func ParseGeoJSON(filename string, ch chan *Result) (*GeoJSON, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := new(GeoJSON)

	if err := json.Unmarshal(data, result); err != nil {
		log.Println(err)
		return nil, err
	}

	ch <- &Result{filename, result}

	return result, nil
}

func ConvertToCountry(g *GeoJSON) *Country {
	c := Country{}
	c.ID = xid.New().String()
	c.Code = g.Features[0].Properties.Adm0Pcode
	c.Name = g.Features[0].Properties.Adm0En
	c.Geometry = g.Features[0].Geometry
	content, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ioutil.WriteFile(countryOut, content, 0644)
	return &c
}

func ConvertToRegions(g *GeoJSON) []Region {
	r := []Region{}
	for _, item := range g.Features {
		region := Region{
			ID:       xid.New().String(),
			Name:     item.Properties.Adm1En,
			Code:     item.Properties.Adm1Pcode,
			Province: item.Properties.OldProvin,
			Geometry: item.Geometry,
		}
		r = append(r, region)
	}
	content, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ioutil.WriteFile(regionsOut, content, 0644)
	return r
}

func ConvertToDistricts(g *GeoJSON) []District {
	d := []District{}
	for _, item := range g.Features {
		district := District{
			ID:       xid.New().String(),
			Name:     item.Properties.Adm2En,
			Code:     item.Properties.Adm2Pcode,
			Province: item.Properties.OldProvin,
			Geometry: item.Geometry,
		}
		d = append(d, district)
	}
	content, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ioutil.WriteFile(districtsOut, content, 0644)
	return d
}

func ConvertToCommunes(g *GeoJSON) []Commune {
	c := []Commune{}
	for _, item := range g.Features {
		commune := Commune{
			ID:       xid.New().String(),
			Name:     item.Properties.Adm3En,
			Code:     item.Properties.Adm3Pcode,
			Province: item.Properties.OldProvin,
			Geometry: item.Geometry,
		}
		c = append(c, commune)
	}
	content, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ioutil.WriteFile(communesOut, content, 0644)
	return c
}

func ConvertToFokontany(g *GeoJSON) []Fokontany {
	f := []Fokontany{}
	for _, item := range g.Features {
		fokontany := Fokontany{
			ID:       xid.New().String(),
			Name:     item.Properties.Adm4En,
			Code:     item.Properties.Adm4Pcode,
			Province: item.Properties.OldProvin,
			Geometry: item.Geometry,
		}
		f = append(f, fokontany)
	}
	content, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ioutil.WriteFile(fokontanyOut, content, 0644)
	return f
}
