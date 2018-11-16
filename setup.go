// +build !nogpu
// +build !linux !arm64

package nvidiasmi

func Init() {
	defer close(initialized)
	HasGPU = false
	info, err := New()
	if err != nil {
		log.WithError(err).Info("was not able to get nvidia-smi info")
		return
	}
	Info = info
	GPUCount = len(Info.GPUS)
	HasGPU = GPUCount > 0
}
