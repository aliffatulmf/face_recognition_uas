package src

import (
	"os/exec"
)

func Python(args ...string) ([]byte, error) {
	cmd := exec.Command("face_recognition/env/Scripts/python.exe", args...)
	return cmd.CombinedOutput()
}

func TrainingModel() string {
	stdout, _ := Python("face_recognition/encode_faces.py",
		"-d", "hog",
		"-i", "face_recognition/dataset",
		"-e", "face_recognition/models/faces.model",
	)

	return string(stdout)
}

func FaceRecognize(input string) string {
	stdout, _ := Python("face_recognition/recognize_faces_image.py",
		"-e", "face_recognition/models/faces.model",
		"-d", "hog",
		"-i", input,
		"-o", input,
	)

	return string(stdout)
}
