package builderService

type Page struct {
  HtmlTop map[string]string
  HtmlBottom map[string]string
  HtmlLeft map[string]string
  HtmlRight map[string]string
}

type BuilderService interface {
  SetDataToHtml() BuilderService
  GetBuild() Page
}

type BuildDirector struct {
  BuildServiceInstance BuilderService
}

func NewBuildDirector() BuildDirector {
  return BuildDirector{}
}

func (bd *BuildDirector) Construct() Page {
  // Construct the View
  bd.BuildServiceInstance = bd.BuildServiceInstance.SetDataToHtml()
  return bd.BuildServiceInstance.GetBuild()
}

func (bd *BuildDirector) SetBuilderService(bs BuilderService) {
  bd.BuildServiceInstance = bs
}

type HomePageBuilderService struct{
  HomePage Page
}

func NewHomePageBuilderService() BuilderService {
  return HomePageBuilderService{}
}

func (hpbs HomePageBuilderService) SetDataToHtml() BuilderService {
  hpbs.HomePage = Page{
    HtmlTop: map[string]string{
      "top":"Home Page Top",
    },
    HtmlBottom: map[string]string{
      "bottom":"Home Page Bottom",
    },
    HtmlLeft: map[string]string{
      "left":"Home Page Left",
    },
    HtmlRight: map[string]string{
      "right":"Home Page Right",
    },
  }
  return hpbs
}

func (hpbs HomePageBuilderService) GetBuild() Page {
  return hpbs.HomePage
}
