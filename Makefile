build:
	fyne package -os android -appID lispoid.mobile -icon lispoid.png -name sync

release:
	fyne package -os android -appID lispoid.mobile -icon lispoid.png -name lispoid --release

debug: build
	adb install lispoid.apk