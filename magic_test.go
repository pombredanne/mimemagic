package mimemagic

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestMatchMagic(t *testing.T) {
	tests := []struct {
		filename string
		want     string
	}{
		{"2001_compression_overview.djvu", "image/vnd.djvu+multipage"},
		{"32x-rom.32x", "application/x-genesis-32x-rom"},
		{"4jsno.669", "audio/x-mod"},
		{"560051.xml", "text/html"},
		{"adf-test.adf", "application/x-amiga-disk-format"},
		{"aero_alt.cur", "image/x-win-bitmap"},
		{"all_w.m3u8", "application/vnd.apple.mpegurl"},
		{"Anaphraseus-1.21-beta.oxt", "application/zip"},
		{"ancp.pcap", "application/vnd.tcpdump.pcap"},
		{"androide.k7", "application/octet-stream"},
		{"aportis.pdb", "application/x-aportisdoc"},
		{"archive.7z", "application/x-7z-compressed"},
		{"archive.lrz", "application/x-lrzip"},
		{"archive.tar", "application/x-tar"},
		{"ascii.stl", "model/stl"},
		{"atari-2600-test.A26", "application/octet-stream"},
		{"atari-7800-test.A78", "application/x-atari-7800-rom"},
		{"atari-lynx-chips-challenge.lnx", "application/x-atari-lynx-rom"},
		{"attachment.tif", "image/jpeg"},
		{"balloon.j2c", "image/x-jp2-codestream"},
		{"balloon.jp2", "image/jp2"},
		{"balloon.jpf", "image/jpx"},
		{"balloon.jpm", "image/jpm"},
		{"balloon.mj2", "video/mj2"},
		{"bathead.sk", "image/x-skencil"},
		{"bbc.ram", "application/vnd.rn-realmedia"},
		{"bibtex.bib", "text/x-bibtex"},
		{"binary.stl", "application/octet-stream"},
		{"blitz.m7", "application/octet-stream"},
		{"bluerect.mdi", "image/vnd.ms-modi"},
		{"bluish.icc", "application/vnd.iccprofile"},
		{"break.mtm", "audio/x-mod"},
		{"bug106330.iso", "application/octet-stream"},
		{"bug-30656-xchat.conf", "text/plain"},
		{"bug39126-broken.ppm", "image/x-portable-pixmap"},
		{"bug39126-working.ppm", "image/x-portable-pixmap"},
		{"build.gradle", "text/plain"},
		{"ccfilm.axv", "video/annodex"},
		{"classiq1.hfe", "application/x-hfe-floppy-image"},
		{"colormapped.tga", "image/x-tga"},
		{"combined.karbon", "application/x-karbon"},
		{"comics.cb7", "application/x-7z-compressed"},
		{"comics.cbt", "application/x-tar"},
		{"COPYING.asc", "application/pgp-signature"},
		{"copying.cab", "application/vnd.ms-cab-compressed"},
		{"COPYING-clearsign.asc", "text/plain"},
		{"COPYING-encrypted.asc", "application/pgp-encrypted"},
		{"core", "application/x-zerosize"},
		{"Core", "application/x-zerosize"},
		{"ct_faac-adts.aac", "audio/aac"},
		{"cube.igs", "model/iges"},
		{"cube.wrl", "model/vrml"},
		{"cyborg.med", "audio/x-mod"},
		{"dbus-comment.service", "text/x-dbus-service"},
		{"dbus.service", "text/x-dbus-service"},
		{"debian-goodies_0.63_all.deb", "application/vnd.debian.binary-package"},
		{"dia.shape", "application/x-dia-shape"},
		{"disk.img", "application/x-zerosize"},
		{"disk.img.xz", "application/x-xz"},
		{"disk.raw-disk-image", "application/x-zerosize"},
		{"disk.raw-disk-image.xz", "application/x-xz"},
		{"dns.cap", "application/vnd.tcpdump.pcap"},
		{"editcopy.png", "image/png"},
		{"Elephants_Dream-360p-Stereo.webm", "video/webm"},
		{"Empty.chrt", "application/x-kchart"},
		{"en_US.zip.meta4", "application/metalink4+xml"},
		{"esm.mjs", "text/plain"},
		{"evolution.eml", "application/mbox"},
		{"example_42_all.snap", "application/vnd.squashfs"},
		{"example.heic", "application/octet-stream"},
		{"example.heif", "application/octet-stream"},
		{"feed2", "application/rss+xml"},
		{"feed.atom", "application/atom+xml"},
		{"feed.rss", "application/rss+xml"},
		{"feeds.opml", "text/x-opml+xml"},
		{"foo-0.1-1.fc18.src.rpm", "application/x-rpm"},
		{"foo.doc", "application/msword"},
		{"fuji.themepack", "application/vnd.ms-cab-compressed"},
		{"game-boy-color-test.gbc", "application/x-gameboy-color-rom"},
		{"game-boy-test.gb", "application/x-gameboy-rom"},
		{"game-gear-test.gg", "application/octet-stream"},
		{"GammaChart.exr", "image/x-exr"},
		{"gedit.flatpakref", "application/vnd.flatpak.ref"},
		{"genesis1.bin", "application/x-genesis-rom"},
		{"genesis2.bin", "application/x-genesis-rom"},
		{"gnome.flatpakrepo", "application/vnd.flatpak.repo"},
		{"good-1-delta-lzma2.tiff.xz", "application/x-xz"},
		{"googleearth.kml", "application/xml"},
		{"gtk-builder.ui", "application/x-gtk-builder"},
		{"hbo-playlist.qtl", "application/x-quicktime-media-link"},
		{"hello.flatpak", "application/vnd.flatpak"},
		{"hello.pack", "application/x-java-pack200"},
		{"helloworld.groovy", "text/x-modelica"},
		{"helloworld.java", "text/x-modelica"},
		{"helloworld.xpi", "application/zip"},
		{"hello.xdgapp", "application/vnd.flatpak"},
		{"hereyes_remake.mo3", "audio/x-mo3"},
		{"html4.css", "text/plain"},
		{"html5.css", "text/x-csrc"},
		{"image.sqsh", "application/vnd.squashfs"},
		{"img_5304.jpg", "image/jpeg"},
		{"internet.ez", "text/plain"},
		{"isdir.m", "text/x-matlab"},
		{"ISOcyr1.ent", "text/html"},
		{"iso-file.iso", "application/octet-stream"},
		{"IWAD.WAD", "application/x-doom-wad"},
		{"javascript-without-extension", "application/javascript"},
		{"jc-win.ani", "application/x-navi-animation"},
		{"json_array.json", "text/plain"},
		{"json-ld-full-iri.jsonld", "text/plain"},
		{"json_object.json", "text/plain"},
		{"layersupdatesignals.flw", "application/x-kivio"},
		{"Leafpad-0.8.17-x86_64.AppImage", "application/x-iso9660-appimage"},
		{"Leafpad-0.8.18.1.glibc2.4-x86_64.AppImage", "application/vnd.appimage"},
		{"libcompat.a", "application/x-archive"},
		{"libcompat.ar", "application/x-archive"},
		{"LiberationSans-Regular.ttf", "font/ttf"},
		{"LiberationSans-Regular.woff", "font/woff"},
		{"linguist.ts", "text/vnd.qt.linguist"},
		{"list", "text/plain"},
		{"live-streaming.m3u", "application/vnd.apple.mpegurl"},
		{"lucid-tab-bg.xcf", "image/x-xcf"},
		{"m64p_test_rom.n64", "application/x-n64-rom"},
		{"m64p_test_rom.v64", "application/x-n64-rom"},
		{"m64p_test_rom.z64", "application/x-n64-rom"},
		{"Makefile", "text/plain"},
		{"Makefile.gnu", "text/plain"},
		{"markdown.md", "text/plain"},
		{"mega-drive-rom.gen", "application/x-genesis-rom"},
		{"menu.ini", "text/plain"},
		{"meson.build", "text/plain"},
		{"meson_options.txt", "text/plain"},
		{"Metroid_japan.fds", "application/x-fds-disk"},
		{"msg0001.gsm", "application/octet-stream"},
		{"msx2-metal-gear.msx", "application/octet-stream"},
		{"msx-penguin-adventure.msx", "application/octet-stream"},
		{"my-data.json-patch", "text/plain"},
		{"mypaint.ora", "image/openraster"},
		{"mysum.m", "text/x-matlab"},
		{"neo-geo-pocket-color-test.ngc", "application/x-neo-geo-pocket-color-rom"},
		{"neo-geo-pocket-test.ngp", "application/x-neo-geo-pocket-rom"},
		{"newtonme.pict", "image/x-pict"},
		{"nrl.trig", "text/plain"},
		{"ocf10-20060911.epub", "application/epub+zip"},
		{"office.doc", "application/msword"},
		{"one-file.tnef", "application/vnd.ms-tnef"},
		{"ooo25876-2.pct", "image/x-pict"},
		{"ooo-6.0.doc", "application/msword"},
		{"ooo-95.doc", "application/msword"},
		{"ooo.doc", "application/msword"},
		{"ooo.rtf", "application/rtf"},
		{"ooo.sdw", "application/vnd.stardivision.writer"},
		{"ooo.stw", "application/vnd.sun.xml.writer.template"},
		//{"ooo.sxw", "application/vnd.sun.xml.writer"},
		{"ooo-test.fodg", "application/xml"},
		{"ooo-test.fodp", "application/xml"},
		{"ooo-test.fods", "application/xml"},
		{"ooo-test.fodt", "application/xml"},
		{"ooo-test.odg", "application/vnd.oasis.opendocument.graphics"},
		{"ooo-test.odp", "application/vnd.oasis.opendocument.presentation"},
		{"ooo-test.ods", "application/vnd.oasis.opendocument.spreadsheet"},
		{"ooo-test.odt", "application/vnd.oasis.opendocument.text"},
		{"ooo.vor", "application/vnd.stardivision.writer"},
		{"ooo-xp.doc", "application/msword"},
		{"Oriental_tattoo_by_daftpunk22.eps", "image/x-eps"},
		{"panasonic_lumix_dmc_fz38_05.rw2", "image/x-panasonic-rw2"},
		{"pdf-not-matlab", "application/pdf"},
		{"petite-ouverture-a-danser.ly", "text/plain"},
		{"pico-rom.bin", "application/x-sega-pico-rom"},
		{"playlist.asx", "audio/x-ms-asx"},
		{"playlist.mrl", "text/x-mrml"},
		{"playlist.wpl", "application/vnd.ms-wpl"},
		{"plugins.qmltypes", "text/x-qml"},
		{"pocket-word.psw", "application/x-pocket-word"},
		{"pom.xml", "text/plain"},
		{"Presentation.kpt", "application/x-kpresenter"},
		{"project.glade", "application/x-glade"},
		{"PWAD.WAD", "application/x-doom-wad"},
		{"pyside.py", "text/x-python"},
		{"raw-mjpeg.mjpeg", "image/jpeg"},
		{"README.pdf", "application/pdf"},
		{"rectangle.qml", "text/x-qml"},
		{"registry-nt.reg", "text/x-ms-regedit"},
		{"registry.reg", "text/x-ms-regedit"},
		{"reStructuredText.rst", "text/plain"},
		{"rgb-reference.ktx", "image/ktx"},
		{"ringtone.ime", "text/x-iMelody"},
		{"ringtone.m4r", "audio/mp4"},
		{"ringtone.mmf", "application/x-smaf"},
		{"ripoux.sap", "application/x-thomson-sap-image"},
		{"sample1.nzb", "application/x-nzb"},
		{"sample2.amr", "audio/AMR"},
		{"sample.docx", "application/zip"},
		{"sample.png.lzma", "application/octet-stream"},
		{"sample.ppsx", "application/zip"},
		{"sample.pptx", "application/zip"},
		{"sample.vsdx", "application/zip"},
		{"sample.xlsx", "application/zip"},
		{"saturn-test.bin", "application/x-saturn-rom"},
		{"SConscript", "text/plain"},
		{"SConscript.buildinfo", "text/plain"},
		{"SConstruct", "text/plain"},
		{"sega-cd-test.iso", "application/x-sega-cd-rom"},
		{"serafettin.rar", "application/vnd.rar"},
		{"settings.xml", "text/plain"},
		{"settopbox.ts", "video/mp2t"},
		{"sg1000-test.sg", "application/octet-stream"},
		{"shebang.qml", "text/x-qml"},
		{"simon.669", "audio/x-mod"},
		{"simple-obj-c.m", "text/x-objcsrc"},
		{"small_wav.mxf", "application/mxf"},
		{"sms-test.sms", "application/octet-stream"},
		{"spinboxes-0.1.1-Linux.tar.xz", "application/x-xz"},
		{"sqlite2.kexi", "application/x-sqlite2"},
		{"sqlite3.kexi", "application/vnd.sqlite3"},
		{"ssh-public-key.txt", "text/plain"},
		{"Stallman_Richard_-_The_GNU_Manifesto.fb2", "application/x-fictionbook+xml"},
		{"Stallman_Richard_-_The_GNU_Manifesto.fb2.zip", "application/x-zip-compressed-fb2"},
		{"stream.nsc", "application/x-netshow-channel"},
		{"stream.sdp", "application/sdp"},
		{"subshapes.swf", "application/vnd.adobe.flash.movie"},
		{"subtitle-microdvd.sub", "text/x-microdvd"},
		{"subtitle-mpsub.sub", "text/x-mpsub"},
		{"subtitle.smi", "application/x-sami"},
		{"subtitle.srt", "application/x-subrip"},
		{"subtitle.ssa", "text/x-ssa"},
		{"subtitle-subviewer.sub", "text/x-subviewer"},
		{"survey.js", "text/x-csrc"},
		{"systemd.automount", "text/x-systemd-unit"},
		{"systemd.device", "text/x-systemd-unit"},
		{"systemd.mount", "text/x-systemd-unit"},
		{"systemd.path", "text/x-systemd-unit"},
		{"systemd.scope", "text/x-systemd-unit"},
		{"systemd.service", "text/x-systemd-unit"},
		{"systemd.slice", "text/x-systemd-unit"},
		{"systemd.socket", "text/x-systemd-unit"},
		{"systemd.swap", "text/x-systemd-unit"},
		{"systemd.target", "text/x-systemd-unit"},
		{"systemd.timer", "text/x-systemd-unit"},
		{"tb-from-sentbox.eml", "message/rfc822"},
		{"tb-saved.eml", "message/rfc822"},
		{"test10.gpx", "application/xml"},
		{"test1.pcf", "text/plain"},
		{"test2.pcf", "application/x-font-pcf"},
		{"test2.ppm", "image/x-portable-pixmap"},
		//{"test2.tga", "image/x-tga"},
		{"test3.py", "text/x-python3"},
		{"test.aa", "audio/x-pn-audibleaudio"},
		{"test.aax", "audio/x-pn-audibleaudio"},
		{"test.aiff", "audio/x-aiff"},
		{"test.alz", "application/x-alz"},
		{"test.avf", "video/x-msvideo"},
		{"test.avi", "video/x-msvideo"},
		{"test.bflng", "text/html"},
		{"test.bmp", "image/bmp"},
		{"test.bsdiff", "application/x-bsdiff"},
		{"testcase.is-really-a-pdf", "application/pdf"},
		{"testcases.ksp", "application/x-kspread"},
		{"test.cbl", "text/plain"},
		{"test.ccmx", "application/x-ccmx"},
		{"test-cdda.toc", "application/x-cdrdao-toc"},
		{"test-cdrom.toc", "application/x-cdrdao-toc"},
		{"test.cel", "application/octet-stream"},
		{"test.cl", "text/x-csrc"},
		{"test.class", "application/x-java"},
		{"test.cmake", "text/plain"},
		{"test.coffee", "text/plain"},
		{"testcompress.z", "application/x-compress"},
		{"test.cs", "text/plain"},
		{"test.csvs", "text/plain"},
		{"test.d", "text/plain"},
		{"test.dcm", "application/dicom"},
		{"test.djvu", "image/vnd.djvu"},
		{"test.dot", "text/vnd.graphviz"},
		{"test.dts", "audio/vnd.dts"},
		{"test.dtshd", "audio/vnd.dts.hd"},
		{"test-en.mo", "application/x-gettext-translation"},
		{"test-en.po", "text/plain"},
		{"test.eps", "image/x-eps"},
		{"test.ext,v", "text/plain"},
		{"test.feature", "text/plain"},
		{"test.fit", "image/fits"},
		{"test.fl", "application/x-fluid"},
		{"test.flac", "audio/flac"},
		{"test.fli", "application/octet-stream"},
		{"test.gbr", "image/x-gimp-gbr"},
		{"test.gcode", "text/plain"},
		{"test.geo.json", "text/plain"},
		{"test.geojson", "text/plain"},
		{"test-gettext.c", "text/x-csrc"},
		{"test.gif", "image/gif"},
		{"test.gih", "application/octet-stream"},
		{"test.gnd", "application/gnunet-directory"},
		{"test.go", "text/plain"},
		{"test.gpx", "application/xml"},
		{"test.gs", "text/x-modelica"},
		{"test.h5", "application/x-hdf"},
		{"test.hdf4", "application/x-hdf"},
		{"test.hlp", "application/winhlp"},
		{"test.ico", "image/vnd.microsoft.icon"},
		{"test.ilbm", "image/x-ilbm"},
		{"test.im1", "image/x-sun-raster"},
		{"test.iptables", "text/x-iptables"},
		{"test.ipynb", "application/x-ipynb+json"},
		{"test.it87", "application/x-it87"},
		{"test.jar", "application/zip"},
		{"test.jceks", "application/x-java-jce-keystore"},
		{"test.jks", "application/x-java-keystore"},
		{"test.jnlp", "application/x-java-jnlp-file"},
		{"test.jpg", "image/jpeg"},
		{"test.kdc", "image/x-kodak-kdc"},
		{"test.key", "application/x-iwork-keynote-sffkey"},
		{"test-kounavail2.kwd", "application/x-kword"},
		{"test.lwp", "application/vnd.lotus-wordpro"},
		{"test.lz", "application/x-lzip"},
		{"test.lz4", "application/x-lz4"},
		{"test.lzo", "application/x-lzop"},
		{"test.manifest", "text/cache-manifest"},
		{"test.metalink", "application/metalink+xml"},
		{"test.mml", "text/plain"},
		{"test.mng", "video/x-mng"},
		{"test.mo", "text/x-modelica"},
		{"test.mobi", "application/x-mobipocket-ebook"},
		{"test.mof", "text/x-csrc"},
		{"test.msi", "application/x-ole-storage"},
		{"test-noid3.mp3", "audio/mpeg"},
		{"test.ogg", "audio/x-vorbis+ogg"},
		{"test.ooc", "text/plain"},
		{"test.opus", "audio/x-opus+ogg"},
		{"test.owx", "application/owl+xml"},
		{"test.p12", "application/octet-stream"},
		{"test-p6.ppm", "image/x-portable-pixmap"},
		{"test.p7b", "application/octet-stream"},
		{"test.pat", "image/x-gimp-pat"},
		{"test.pbm", "image/x-portable-bitmap"},
		{"test.pcx", "image/vnd.zbrush.pcx"},
		{"test.pdf.lz", "application/x-lzip"},
		{"test.pdf.xz", "application/x-xz"},
		{"test.pgm", "image/x-portable-graymap"},
		{"test.pgn", "application/vnd.chess-pgn"},
		{"test.php", "application/x-php"},
		{"test.pix", "application/octet-stream"},
		{"test.pkipath", "application/octet-stream"},
		{"test.pl", "application/x-perl"},
		{"test.pm", "application/x-perl"},
		{"test.pmd", "application/x-ole-storage"},
		{"test.png", "image/png"},
		{"test.por", "application/x-spss-por"},
		{"test.pot", "text/x-gettext-translation-template"},
		{"test.ppm", "image/x-portable-pixmap"},
		{"test.ps", "application/postscript"},
		{"test.psd", "image/vnd.adobe.photoshop"},
		{"test-public-key.asc", "application/pgp-keys"},
		{"test.py", "text/x-python"},
		{"test.py3", "text/x-python3"},
		{"test.pyx", "text/plain"},
		{"test.qp", "application/x-qpress"},
		{"test.qti", "application/x-qtiplot"},
		{"test.raml", "application/raml+yaml"},
		{"test.random", "application/octet-stream"},
		{"test-reordered.ipynb", "text/plain"},
		{"test.rs", "text/plain"},
		{"test.sass", "text/plain"},
		{"test.sav", "application/x-spss-sav"},
		{"test.scala", "text/plain"},
		{"test.scm", "text/plain"},
		{"test.scss", "text/plain"},
		{"test-secret-key.asc", "application/pgp-keys"},
		{"test-secret-key.skr", "application/pgp-keys"},
		{"test.sgi", "application/octet-stream"},
		{"test.sqlite2", "application/x-sqlite2"},
		{"test.sqlite3", "application/vnd.sqlite3"},
		{"test.ss", "text/plain"},
		{"test.sv", "text/plain"},
		{"test.svh", "text/plain"},
		{"test.t", "application/x-perl"},
		{"test.tar.lz", "application/x-lzip"},
		{"test.tar.lz4", "application/x-lz4"},
		{"test-template.dot", "application/msword"},
		{"test.tex", "text/x-tex"},
		{"test.tga", "image/x-tga"},
		{"test.tif", "image/tiff"},
		{"test.ts", "video/mp2t"},
		{"test.ttl", "text/plain"},
		{"test.ttx", "application/x-font-ttx"},
		{"test.twig", "text/plain"},
		{"test.url", "application/x-mswinurl"},
		{"test.uue", "text/x-uuencode"},
		{"test.v", "text/plain"},
		{"test.vala", "text/plain"},
		{"test.vcf", "text/vcard"},
		{"test-vpn.pcf", "application/x-cisco-vpn-settings"},
		{"test.vsd", "application/x-ole-storage"},
		{"test.wav", "audio/x-wav"},
		{"test.webp", "image/webp"},
		{"test.wim", "application/x-ms-wim"},
		{"test.wps", "application/x-ole-storage"},
		{"test.xar", "application/x-xar"},
		{"test.xbm", "text/plain"},
		{"test.xcf", "image/x-xcf"},
		{"test.xht", "application/xhtml+xml"},
		{"test.xhtml", "application/xhtml+xml"},
		{"test.xlr", "application/x-ole-storage"},
		{"test.xml.in", "application/xml"},
		{"test.xpm", "image/x-xpixmap"},
		{"test.xsl", "application/xslt+xml"},
		{"test.xwd", "application/octet-stream"},
		{"test.yaml", "application/x-yaml"},
		{"test.zip", "application/zip"},
		{"test.zz", "application/octet-stream"},
		{"text-iso8859-15.txt", "text/plain"},
		{"text.pdf", "application/pdf"},
		{"text.ps", "application/postscript"},
		{"text.ps.gz", "application/gzip"},
		{"text.PS.gz", "application/gzip"},
		{"text.qmlproject", "text/x-qml"},
		{"text-utf8.txt", "text/plain"},
		{"text.wwf", "application/pdf"},
		{"tree-list", "text/plain"},
		{"TS010082249.pub", "application/x-ole-storage"},
		{"upc-video-subtitles-en.vtt", "text/vtt"},
		{"Utils.jsm", "application/x-perl"},
		{"virtual-boy-wario-land.vb", "application/octet-stream"},
		{"weather_sun.xcf", "image/x-xcf"},
		{"webfinger.jrd", "text/plain"},
		{"white_640x480.kra", "application/x-krita"},
		{"wii.wad", "application/x-wii-wad"},
		{"wonderswan-color-chocobo.wsc", "application/octet-stream"},
		{"wonderswan-rockman-forte.ws", "application/octet-stream"},
		{"xml-in-mp3.mp3", "audio/mpeg"},
	}
	path, err := unpackFixtures()
	if err != nil {
		t.Fatalf("couldn't unpack archive: %v", err)
	}
	defer func() {
		err := os.RemoveAll(path)
		if err != nil {
			panic(err)
		}
	}()
	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			data, err := ioutil.ReadFile(filepath.Join(path, test.filename))
			if err != nil {
				t.Fatalf("couldn't read file %s: %v", test.filename, err)
			}
			if len(data) > magicMaxLen {
				data = data[:magicMaxLen]
			}
			if got := MatchMagic(data).MediaType(); got != test.want {
				t.Errorf("MatchMagic() = %v, want %v", got, test.want)
			}
		})
	}
	t.Run("BOM", func(t *testing.T) {
		data := []byte{0xff, 0xfe, 0x00}
		want := "text/plain"
		if got := MatchMagic(data).MediaType(); got != want {
			t.Errorf("MatchMagic() = %v, want %v", got, want)
		}
	})
}

func benchmarkMatchMagic(filename string, b *testing.B) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		b.Fatalf("couldn't read file %s: %v", filename, err)
	}
	if len(data) > magicMaxLen {
		data = data[:magicMaxLen]
	}
	for n := 0; n < b.N; n++ {
		MatchMagic(data)
	}
}

func BenchmarkMatchMagic(b *testing.B) {
	path, err := unpackFixtures()
	if err != nil {
		b.Fatalf("couldn't unpack archive: %v", err)
	}
	defer func() {
		err := os.RemoveAll(path)
		if err != nil {
			panic(err)
		}
	}()
	for _, f := range combinedTests {
		b.Run(f.filename, func(b *testing.B) {
			benchmarkMatchMagic(filepath.Join(path, f.filename), b)
		})
	}
}
