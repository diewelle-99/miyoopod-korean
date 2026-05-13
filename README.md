# MiyooPod Korean Patch

MiyooPod Korean Patch는 [Danilo Fragoso](https://github.com/danfragoso)의 원본 MiyooPod를 기반으로 한 한국어 수정판입니다.

원본 MiyooPod는 클래식 iPod 스타일 인터페이스에서 영감을 받은 Miyoo용 음악 플레이어입니다. 이 수정판은 한국어 UI, 한국어 폰트, 초성 검색, 재생 진행률 보정, 곡 길이 보정, Now Playing 화면의 트랙별 앨범아트 표시 개선을 목표로 합니다.

![MiyooPod](screenshots/hero.png)

> **현재 상태**
>
> - 한국어 UI/폰트/검색 패치 적용
> - Ubuntu 20.04 WSL + `arm-linux-gnueabihf-gcc` 환경에서 ARM 32-bit 빌드 확인
> - `GLIBC_2.4` 기준 빌드 확인


## 주요 기능

- iPod에서 영감을 받은 사용자 인터페이스와 여러 테마
- 아티스트, 앨범, 노래 기준 탐색
- 화면 키보드를 이용한 목록 검색/필터
- 한국어 UI 문구
- 한국어 폰트 적용
- 한글 초성 검색 지원
- MusicBrainz를 통한 앨범아트 자동 다운로드
- 내장 앨범아트 표시
- Now Playing 화면의 트랙별 앨범아트 표시 개선
- 셔플 및 반복 재생 모드
- 가속 방식의 탐색, 빨리감기, 되감기
- OTA 업데이트 기능
- 세션 유지 기능: 대기열, 재생 위치, 셔플/반복 상태 복원
- Miyoo Mini 계열에 최적화된 640×480 해상도 UI
- 17개 사용자 지정 테마
- MP3, FLAC, OGG/Vorbis 재생 지원
- LRC 시간 동기화 가사 및 자동 스크롤 가사 표시
- SDL_mixer가 재생 위치를 제대로 반환하지 않는 환경을 위한 재생 진행률 보정
- 일부 파일에서 곡 길이가 0:00으로 표시되는 문제를 줄이기 위한 길이 추정 보정

## 설치 방법

1. GitHub Releases에서 `MiyooPod-Korean.zip`을 다운로드합니다.
2. ZIP 파일을 압축 해제합니다.
3. Miyoo Mini 또는 Miyoo Mini Plus의 SD카드를 컴퓨터에 연결합니다.
4. 압축 해제한 `MiyooPod` 폴더를 SD카드의 `/App` 폴더에 복사합니다.
5. SD카드를 안전하게 제거한 뒤 기기에 다시 삽입합니다.
6. Apps 메뉴에서 MiyooPod를 실행합니다.

기본 설치 경로는 다음과 같습니다.

```text
/mnt/SDCARD/App/MiyooPod/
```

SD카드에서는 보통 아래 구조가 되어야 합니다.

```text
/App/MiyooPod/
├─ launch.sh
├─ MiyooPod
├─ updater
├─ assets/
│  └─ ui_font.ttf
└─ libs/
   ├─ libSDL2.so
   ├─ libSDL2_mixer.so
   └─ 기타 필요한 .so 파일
```

## 음악 추가 방법

MiyooPod는 SD카드의 Music 폴더에서 음악 파일을 읽습니다.

```text
/mnt/SDCARD/Media/Music/
```

1. Miyoo 기기의 SD카드를 컴퓨터에 연결합니다.
2. `/Media/Music/` 폴더로 이동합니다.
3. MP3, FLAC, OGG/Vorbis 파일과 음악 폴더를 복사합니다.
4. 아티스트/앨범 폴더로 정리하면 라이브러리 탐색이 더 편합니다.
5. MiyooPod를 실행하면 음악 라이브러리를 자동으로 스캔하고 색인합니다.

## 지원 형식

MiyooPod는 다음 오디오 파일을 지원합니다.

- **MP3**
- **FLAC**
- **OGG/Vorbis**

**권장 형식:** MP3 @ 256kbps

- **형식:** MP3, MPEG-1 Audio Layer 3
- **비트레이트:** 256kbps CBR 또는 VBR V0
- **샘플레이트:** 44.1kHz
- **채널:** 스테레오

> **참고:** Miyoo Mini 계열의 오디오 출력 품질과 하드웨어 성능을 고려하면 고비트레이트 파일이 큰 이점을 주지 않을 수 있습니다. 너무 높은 비트레이트의 파일은 재생이 끊길 수 있습니다. FLAC도 지원하지만, 성능과 안정성 면에서는 MP3를 권장합니다.

## 앨범아트

### 내장 앨범아트

MiyooPod는 MP3 파일의 ID3 태그에 포함된 앨범아트를 자동으로 추출합니다.

한국어 수정판에서는 Now Playing 화면에서 같은 앨범 안의 곡이라도 파일마다 다른 내장 앨범아트를 표시할 수 있도록 트랙별 앨범아트 표시를 개선했습니다.

### MusicBrainz 자동 다운로드

내장 앨범아트가 없는 앨범은 MusicBrainz를 통해 앨범 커버를 자동으로 가져올 수 있습니다.

1. 메인 메뉴에서 **설정**으로 이동합니다.
2. **앨범아트 가져오기**를 선택합니다.
3. MiyooPod가 라이브러리를 스캔하고 누락된 앨범아트를 다운로드합니다.

> **참고:** Wi-Fi 인터넷 연결이 필요합니다. 앨범아트는 다음 폴더에 저장됩니다.
>
> ```text
> /mnt/SDCARD/Media/Music/.miyoopod_artwork/
> ```

## 설정

- **테마** - Classic iPod, Dark, Dark Blue, Light, Nord, Solarized Dark, Matrix Green, Retro Amber, Purple Haze, Cyberpunk, Coffee, Ocean, Forest, Sunset, Neon, Midnight, Gruvbox, Candy 등 17개 테마 선택
- **잠금 키** - 화면 잠금/해제 버튼 설정: Y, X, SELECT
- **앨범아트 가져오기** - MusicBrainz에서 누락된 앨범아트를 자동 다운로드
- **업데이트 확인** - OTA 업데이트를 수동으로 확인하고 설치
- **업데이트 알림** - 자동 업데이트 알림 켜기/끄기
- **앱 데이터 초기화** - 라이브러리 캐시, 설정, 앨범아트 초기화
- **로그 토글** - 디버그 로그 켜기/끄기
- **라이브러리 다시 스캔** - 음악 라이브러리 전체 재스캔
- **정보** - 앱 버전 확인 및 업데이트 확인

## 한국어 패치 내용

이 수정판에는 다음 변경사항이 포함되어 있습니다.

- 한국어 표시가 가능한 `ui_font.ttf` 적용
- 주요 메뉴와 상태 문구 한국어화
- 검색 화면에 한글 초성 입력 추가
- 한글 문자열 삭제/입력 시 깨짐을 줄이기 위한 rune 기반 처리
- 아티스트, 앨범, 곡 제목 검색 시 초성 검색 지원
- SDL_mixer가 재생 위치를 반환하지 않는 환경에서 재생 시간이 0:00에 멈추는 문제 보정
- 일부 MP3/FLAC 파일의 전체 길이가 0:00으로 표시되는 문제를 줄이기 위한 길이 추정 보정
- 라이브러리 스캔 중 첫 파일 처리에서 0곡 상태로 오래 멈춰 보이는 문제 완화
- Now Playing 화면에서 트랙별 내장 앨범아트 우선 표시
- 곡 변경 시 Now Playing 캐시 갱신

## 기술 정보

Go 1.22.2와 CGO 기반 C 바인딩을 사용하여 그래픽과 오디오를 처리합니다.

### 아키텍처

- **플랫폼:** 원본 기준 Miyoo Mini Plus, Mini v4, Mini Flip / OnionOS
- **현재 한국어 수정판 빌드 기준:** ARM 32-bit, Miyoo Mini / Mini Plus 계열 테스트용
- **CPU:** ARM Cortex-A7 계열
- **해상도:** 원본 기준 640×480 native, 일부 기기에서 자동 해상도 감지
- **크로스 컴파일:** `arm-linux-gnueabihf-gcc`
- **확인된 GLIBC 요구사항:** `GLIBC_2.4`

> **Miyoo Flip / SpruceUI 참고**
>
> 이 저장소의 현재 한국어 수정판은 Flip/SpruceUI 전용 포팅이 완료된 상태가 아닙니다. Flip/SpruceUI에서 사용하려면 실행 스크립트, SDL2 라이브러리, 입력 매핑, 오디오/비디오 드라이버 확인이 별도로 필요합니다.

### 주요 라이브러리

- **SDL2** - 그래픽, 입력 처리, 창 관리
- **SDL2_mixer** - MP3 디코딩을 포함한 오디오 재생
- **fogleman/gg** - 2D 그래픽 렌더링
- **dhowden/tag** - ID3 태그 파싱
- **golang.org/x/image** - 이미지 처리 및 폰트 렌더링

### 성능 최적화

- UI와 오디오 처리를 분리한 듀얼코어 활용
- 시간 표시용 숫자 스프라이트 사전 렌더링
- 텍스트 측정 결과와 앨범아트 캐시
- 빠른 시작을 위한 JSON 기반 라이브러리 메타데이터 캐시
- Now Playing 화면의 불필요한 전체 redraw 최소화

## 문제 해결

### 앱이 실행되지 않거나 시작 시 튕기는 경우

라이브러리 캐시 파일이 손상되었을 수 있습니다. SD카드를 컴퓨터에 연결하고 음악 폴더의 숨김 JSON 파일을 삭제하세요.

```text
/mnt/SDCARD/Media/Music/.miyoopod_library.json
/mnt/SDCARD/Media/Music/.miyoopod_state.json
```

다음 실행 시 MiyooPod가 라이브러리를 다시 스캔합니다.

### 한국어가 네모칸으로 보이는 경우

`assets/ui_font.ttf`가 한국어 글리프를 포함한 폰트인지 확인하세요. 한국어 수정판에서는 Noto Sans KR 계열 폰트를 사용하는 것을 권장합니다.

확인할 경로:

```text
/mnt/SDCARD/App/MiyooPod/assets/ui_font.ttf
```

### 재생바와 시간이 0:00에 멈춰 있는 경우

일부 SDL_mixer 환경에서는 재생 위치 또는 전체 길이를 제대로 반환하지 못할 수 있습니다. 한국어 수정판에는 재생 위치 fallback과 MP3/FLAC 길이 추정 fallback이 포함되어 있습니다.

문제가 계속되면 라이브러리 캐시를 삭제한 뒤 다시 스캔하세요.

```text
/mnt/SDCARD/Media/Music/.miyoopod_library.json
/mnt/SDCARD/Media/Music/.miyoopod_state.json
```

### 앨범아트가 올바르게 표시되지 않는 경우

**설정** → **앱 데이터 초기화**를 실행하여 앨범아트 캐시와 라이브러리 메타데이터를 초기화하세요. 다음 실행 시 다시 생성됩니다.

수동으로 삭제하려면:

```text
/mnt/SDCARD/Media/Music/.miyoopod_artwork/
```

### 로그 확인

앱이 이상하게 동작하면 **설정** → **로그 토글**을 켠 뒤 문제를 재현하고 SD카드에서 로그 파일을 확인하세요.

```text
/mnt/SDCARD/App/MiyooPod/miyoopod.log
```

### OTA 업데이트 후 앱이 망가진 경우

업데이트 후 앱이 실행되지 않는다면 최신 버전을 수동으로 다시 설치하세요.

1. GitHub Releases에서 최신 ZIP 파일을 다운로드합니다.
2. 압축을 풀고 `MiyooPod` 폴더를 SD카드의 `/App` 폴더에 덮어씁니다.
3. 음악 라이브러리와 설정은 별도 파일에 저장되므로 보통 음악 파일에는 영향이 없습니다.

## 소스에서 빌드하기

Ubuntu 20.04 WSL 기준 예시입니다.

```bash
sudo apt update
sudo apt install -y git make zip build-essential \
  gcc-arm-linux-gnueabihf g++-arm-linux-gnueabihf \
  pkg-config binutils-arm-linux-gnueabihf
```

Go 1.22.2 설치 후:

```bash
cd ~/miyoopod
go mod download
make go
make updater
```

빌드 결과 확인:

```bash
file App/MiyooPod/MiyooPod
arm-linux-gnueabihf-readelf -V App/MiyooPod/MiyooPod | grep GLIBC | sort -u
```

권장 결과:

```text
ELF 32-bit LSB executable, ARM
GLIBC_2.4
```

패키징:

```bash
cd ~/miyoopod/App
zip -r ~/MiyooPod-Korean.zip MiyooPod
```

## Changelog

### Korean Patch / 한국어 수정판

- 🇰🇷 한국어 UI 문구 추가
- 🇰🇷 한국어 표시용 폰트 적용
- 🔍 한글 초성 검색 지원
- ⌨️ 검색 화면에 한글 초성 입력 추가
- 🧭 한글 검색어 삭제/입력 시 문자열 깨짐을 줄이기 위한 rune 기반 처리
- ⏱️ SDL_mixer가 재생 위치를 반환하지 않는 환경에서 재생 진행률 fallback 추가
- ⏳ 일부 MP3/FLAC 파일에서 전체 길이가 0:00으로 표시되는 문제를 줄이기 위한 길이 추정 fallback 추가
- 🔄 라이브러리 스캔 중 첫 파일 처리에서 0곡 상태로 멈춰 보이는 문제 완화
- 🖼️ Now Playing 화면에서 트랙별 내장 앨범아트를 우선 표시하도록 개선
- 🖼️ 곡 변경 시 Now Playing 화면 캐시를 갱신하도록 수정
- 🧪 Ubuntu 20.04 WSL + `arm-linux-gnueabihf-gcc` 환경에서 ARM 32-bit 빌드 확인
- 🧪 `GLIBC_2.4` 기준 빌드 확인

### Version 0.0.6

- 🎵 FLAC 및 OGG/Vorbis 재생 지원: SDL2_mixer 내부의 정적 링크 drflac, stb_vorbis를 통해 디코딩
- 📝 가사 지원: 내장 가사, ID3 USLT, Vorbis comments를 줄바꿈과 스크롤로 표시
- 🎤 LRC 시간 동기화 가사 지원: 현재 줄 하이라이트, 자동 스크롤, 수동 스크롤 오버라이드
- ⏩ ↑/↓ 버튼을 길게 눌러 목록을 연속 스크롤
- ❌ SELECT + START로 앱 종료

### Version 0.0.5

- 🔄 다운로드 진행률, 체크섬 검증, 실패 시 자동 롤백을 포함한 OTA 업데이트
- 🔍 검색: 아티스트, 앨범, 노래 목록을 화면 A-Z 키보드로 필터링, 목록에서 SELECT로 실행
- ⏩ Now Playing 화면에서 L 또는 R을 길게 눌러 가속 방식으로 탐색/빨리감기/되감기
- 💾 세션 유지: 대기열, 재생 위치, 셔플/반복 상태, 현재 트랙을 앱 재실행 후 복원
- 📜 헤더 마키: 메뉴 탐색 중 헤더 바에 현재 재생 정보 스크롤 표시
- 🛡️ 크래시 리포팅: fatal panic 및 C-level signal을 자동으로 기록하고 보고
- 🗑️ 설정에 앱 데이터 초기화 옵션 추가: 라이브러리 캐시, 설정, 앨범아트 초기화
- ⚡ 시작 속도 개선: 버전 확인이 스플래시 화면을 막지 않도록 수정, 앨범아트는 디스크의 빠른 RGBA 픽셀 캐시 사용
- 🔄 전용 진행 화면에서 트랙 수, 현재 폴더, 단계 표시가 가능한 non-blocking 라이브러리 스캔
- 🖼️ 진행률, 퍼센트, 취소/재시도 지원이 있는 non-blocking 앨범아트 가져오기
- 🔊 볼륨과 밝기를 앱 실행 간 유지
- 🖼️ 시작 후 백그라운드에서 MP3 태그의 앨범아트 추출
- 🔔 설정에서 업데이트 알림 켜기/끄기
- 🔍 설정에서 수동 업데이트 확인 옵션 추가
- 🐛 백그라운드 goroutine이 framebuffer를 손상시켜 panic을 유발하던 race condition 수정
- 🐛 부분 framebuffer 업데이트로 인해 볼륨/밝기 오버레이 화면이 깜빡이던 문제 수정
- 🐛 앱 실행 때마다 볼륨이 초기화되던 문제 수정

### Version 0.0.4

- 🔊 Onion/keymon과 일치하는 올바른 indirect buffer layout으로 MI_AO ioctl 볼륨 제어 수정
- 🔊 오버레이에서 볼륨 아이콘 SVG가 잘리던 문제 수정
- 🔒 전원 버튼을 이용한 화면 잠금 추가
- 🔒 자동 화면 잠금 설정 추가: 1/3/5/10분 또는 비활성화
- 🔒 잠금 중 버튼 입력 시 화면 깨우기 여부를 설정하는 screen peek 토글 추가
- 🐛 화면 잠금 중 밝기와 볼륨이 조절되던 문제 수정
- 🐛 Now Playing 진행바가 잠금 오버레이 위에 그려지던 문제 수정

### Version 0.0.3

- 🔧 PostHog 로깅 초기화 순서 수정
- 📊 SDL 초기화의 C 로그가 제대로 캡처되도록 수정
- 📱 기기 모델 감지 및 보고: Mini Plus, Mini v4, Mini Flip
- 📏 디스플레이 해상도 메트릭을 analytics로 전송
- 🔀 로컬 로그와 개발자 로그 설정을 독립적으로 분리

### Version 0.0.2

- ✨ Miyoo Mini v4 지원 추가: 750×560 해상도
- ✨ 원본 기준 Miyoo Mini Flip 지원 항목 추가: 750×560 해상도
- 🔧 framebuffer device를 통한 자동 해상도 감지
- 🎨 화면 크기에 따라 종횡비를 유지하며 UI scaling 적용
- 🐛 로컬 로그를 기본 비활성화, 개발자 로그는 유지

### Version 0.0.1

- 🎉 최초 릴리스
- 🎵 iPod에서 영감을 받은 사용자 인터페이스
- 🎨 11개 사용자 지정 테마
- 🖼️ 앨범아트 표시 및 MusicBrainz 자동 가져오기
- 🔀 셔플 및 반복 모드
- 📱 Miyoo Mini Plus 640×480에 최적화

## 기여

MiyooPod 원본은 오픈소스 프로젝트입니다. 버그 제보, 기능 요청, Pull Request는 원본 저장소 또는 사용자님 수정판 저장소의 Issues/PR 기능을 사용할 수 있습니다.

- **원본 버그 제보:** [GitHub Issues](https://github.com/danfragoso/miyoopod/issues)
- **원본 기능 요청:** [New Issue](https://github.com/danfragoso/miyoopod/issues/new)
- **원본 PR:** [Pull Requests](https://github.com/danfragoso/miyoopod/pulls)

## 라이선스

Open Source

원본 프로젝트의 저작권 및 크레딧은 원저자에게 있습니다. 이 저장소는 원본 MiyooPod를 기반으로 한 한국어 수정판입니다.

## 원저자

Created by [Danilo Fragoso](https://github.com/danfragoso)

## 한국어 수정판

Korean localization and fixes by [diewelle-99](https://github.com/diewelle-99)

---
