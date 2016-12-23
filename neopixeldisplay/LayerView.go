package neopixeldisplay

//each layer has a ColorCombineMode. When rendering we process the layers bottom up, using the CombineMode of the upper layer to combine with the result of all lower layers
//in order to have layers that combine with each other in different ways, use nested LayerViews

type LayerView struct {
	target *ColorFrame
	layers []*LayerData
}

type LayerData struct {
	Frame *ColorFrame
	CombineMode ColorCombineMode
}

func NewLayerView(target *ColorFrame) *LayerView {
	return &LayerView{target, nil}
}

func (lv *LayerView) AddLayer(combineMode ColorCombineMode) *LayerData {
	frame := MakeColorFrame(lv.target.Width, lv.target.Height, MakeColor(0,0,0))
	frame.parent = lv

	data := &LayerData{&frame, combineMode}
	lv.layers = append(lv.layers, data)
	return data
}

//returns true if the layer was found and deleted, false otherwise
func (lv *LayerView) DeleteLayer(layer *LayerData) bool {
	for i := range lv.layers {
		if lv.layers[i] == layer {
			lv.layers = append(lv.layers[:i], lv.layers[i+1:]...)
			lv.Draw()
			return true
		}
	}
	return false
}

func (lv *LayerView) DeleteLayerIndex(index int) bool {
	if index >= 0 && index < len(lv.layers) {
		lv.layers = append(lv.layers[:index], lv.layers[index+1:]...)
		lv.Draw()
		return true
	}
	return false
}

func (lv *LayerView) GetLayer(index int) *LayerData {
	return lv.layers[index]
}

func (lv *LayerView) Draw() {
	lv.target.SetAll(MakeColor(0,0,0))
	for _, layer := range lv.layers {
		lv.target.CombineRect(0,0, *layer.Frame, layer.CombineMode, Error)
	}
	lv.target.Draw()
}
