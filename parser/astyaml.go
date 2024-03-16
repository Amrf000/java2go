package parser

type AstClass struct {
	Root struct {
		Type               string `json:"-type"`
		PackageDeclaration struct {
			Type string `json:"-type"`
			Name struct {
				Type       string `json:"-type"`
				Identifier string `json:"-identifier"`
			} `json:"name"`
		} `json:"packageDeclaration"`
		Types struct {
			Type []struct {
				Name struct {
					Type       string `json:"-type"`
					Identifier string `json:"-identifier"`
				} `json:"name"`
				Comment struct {
					Type    string `json:"-type"`
					Content string `json:"-content"`
				} `json:"comment,omitempty"`
				Members struct {
					Member []struct {
						Type      string `json:"-type"`
						Modifiers struct {
							Modifier struct {
								Type    string `json:"-type"`
								Keyword string `json:"-keyword"`
							} `json:"modifier"`
						} `json:"modifiers"`
						Variables struct {
							Variable struct {
								Name struct {
									Type       string `json:"-type"`
									Identifier string `json:"-identifier"`
								} `json:"name"`
								Type struct {
									Type interface{} `json:"-type"`
									Name struct {
										Type       string `json:"-type"`
										Identifier string `json:"-identifier"`
									} `json:"name,omitempty"`
								} `json:"type"`
								Type1 string `json:"-type"`
							} `json:"variable"`
						} `json:"variables,omitempty"`
						Body struct {
							Statements struct {
								Statement interface{} `json:"statement"`
							} `json:"statements"`
							Type string `json:"-type"`
						} `json:"body,omitempty"`
						Name struct {
							Type       string `json:"-type"`
							Identifier string `json:"-identifier"`
						} `json:"name,omitempty"`
						Parameters struct {
							Parameter interface{} `json:"parameter"`
						} `json:"parameters,omitempty"`
						Type1 struct {
							Type interface{} `json:"-type"`
							Name struct {
								Type       string `json:"-type"`
								Identifier string `json:"-identifier"`
							} `json:"name,omitempty"`
						} `json:"type,omitempty"`
						Comment struct {
							Type    string `json:"-type"`
							Content string `json:"-content"`
						} `json:"comment,omitempty"`
					} `json:"member"`
				} `json:"members"`
				Modifiers struct {
					Modifier struct {
						Type    string `json:"-type"`
						Keyword string `json:"-keyword"`
					} `json:"modifier"`
				} `json:"modifiers"`
				Type          string `json:"-type"`
				IsInterface   string `json:"-isInterface"`
				ExtendedTypes struct {
					ExtendedType struct {
						Type string `json:"-type"`
						Name struct {
							Type       string `json:"-type"`
							Identifier string `json:"-identifier"`
						} `json:"name"`
					} `json:"extendedType"`
				} `json:"extendedTypes,omitempty"`
			} `json:"type"`
		} `json:"types"`
	} `json:"root"`
}
