package imagetovideo

import "testing"

func TestCanResizeImage(t *testing.T) {
	// given : 유효한 이미지

	// when : 이미지 사이즈 변경 요청

	// then : 이미지 저장 파일 확인
}

func TestCanRequestGenerateVideo(t *testing.T) {
	// given : 유효한 토큰 + 이미지

	// when : 비디오 생성 요청

	// then : 생성 요청 확인(generated_id)
}

func TestCannotRequestGenerateVideoWithNonToken(t *testing.T) {
	// given : 유효하지 않은 토큰 + 이미지

	// when : 비디오 생성 요청

	// then : 토큰 에러 메시지
}

func TestCanGetGenerateVideo(t *testing.T) {
	// given : generated_id

	// when : 비디오 전달 요청

	// then : 비디오 확인
}
