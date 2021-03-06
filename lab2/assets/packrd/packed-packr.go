// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "b463f46fabbf7f99fb3eea0d64e6c4a0"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"07590dc7743bb225eba64b733a381bb4": "1f8b08000000000000ff7c914d6bc2401086eff32b869c12aa9782274f6b32b6a1e9ae4cd682a7b2c4ad2c8d9b90ae14ff7d311fc516716ecbf3eccbbccc7c8e0f4777e84cb0b86d015226a109b5581584950b678801113172fb08a7f9b29d33f5f8904aa3dc16c56cf0aae6e44377ee65e7833dd8eea6e7cdd14e896f82d367c1f1e36291fcf7daa63dd526b8c64777f32ad3ba60ea3e72a5544142def2369cbf0adee10bed30be944a66d08354c952b3c8a5c68fcff7b1054c8dd78a297f92c3b71126bff4324c6b6292299538728cddfeafa324665490262c69580a9225c0f50db2e6db0364ac36573758c24f000000ffff7b68d9f0a7010000",
		"357e5313fd79a77d82b39c9583ea9c0f": "1f8b08000000000000ffd2d555d0cecd4c2f4a2c4955082de0e2720e72750c7155087174f2715548ce2fcd2b29aae4d2e05250505050ca4c51525050284e2dca4ccc5180003fff1005bf501f1f1d888abcc4dc54258530c720670fc7200d2353534d341501419ebe8e41910adeae910a1a200335b934adb9b8905de1925f9ec7c5e512e41f80ea0a6b2e2e40000000ffffcf5f7fbdad000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("migrations", "./migrations")
		b.SetResolver("001_country.sql", packr.Pointer{ForwardBox: gk, ForwardPath: "357e5313fd79a77d82b39c9583ea9c0f"})
		b.SetResolver("002_city.sql", packr.Pointer{ForwardBox: gk, ForwardPath: "07590dc7743bb225eba64b733a381bb4"})
	}()
	return nil
}()
