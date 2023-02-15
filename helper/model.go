package helper

import (
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/entity/web"
)

func ToCreateAkunResponse(user domain.User) web.CreateAkunResponse {
	return web.CreateAkunResponse{
		IdUser: user.IdUser,
		Email:  user.Email,
		Status: user.StatusLogin,
	}
}

func ToOtpResponse(otp domain.Otp) web.OtpResponse {
	return web.OtpResponse{
		Email:  otp.Email,
		NoHp:   otp.NoHp,
		Status: "OTP Terkirim",
	}
}

func ToOtpValidationResponse(otp domain.Otp) web.OtpValidationResponse {
	return web.OtpValidationResponse{
		Email:  otp.Email,
		Status: otp.Status,
	}
}

func ToUpdatePasswordResponse(user domain.User) web.UpdatePasswordResponse {
	return web.UpdatePasswordResponse{
		Email:  user.Email,
		Status: "succes",
	}
}

func ToLoginResponse(nasabah domain.User) web.LoginResponse {
	return web.LoginResponse{
		IdUser: nasabah.IdUser,
		Email:  nasabah.Email,
		Status: nasabah.StatusLogin,
	}
}

func ToGetProfileResponse(user domain.User) web.GetProfileResponse {
	return web.GetProfileResponse{
		IdUser: user.IdUser,
		Email:  user.Email,
		Nama:   user.Nama,
		NoHp:   user.NoHp,
		Avatar: user.Avatar,
	}
}

func ToGetEmailResponse(string2 string) web.GetEmailCheckResponse {
	return web.GetEmailCheckResponse{
		Status: string2,
	}
}

func ToGetKaryawanResponse(karyawan domain.Karyawan) web.GetKaryawanResponse {
	return web.GetKaryawanResponse{
		IdKaryawan:  karyawan.IdKaryawan,
		IdJabatan:   karyawan.IdJabatan,
		NamaLengkap: karyawan.NamaLengkap,
		Foto:        karyawan.Foto,
		Alamat:      karyawan.Alamat,
		Agama:       karyawan.Agama,
		Email:       karyawan.Email,
		NoHp:        karyawan.NoHp,
		Pendidikan:  karyawan.Pendidikan,
	}
}

func ToGetOfficeDataResponse(office domain.Office) web.OfficeResponse {
	return web.OfficeResponse{
		NamaKantor: office.NamaKantor,
		Alamat:     office.Alamat,
		Latitude:   office.Latitude,
		Longintude: office.Longitude,
		JamMasuk:   office.JamMasuk,
		JamPulang:  office.JamPulang,
		Radius:     office.Radius,
	}
}

func ToPresensiMasukResponse(presensi domain.Presensi) web.PresensiMasukResponse {
	return web.PresensiMasukResponse{
		Keterangan: presensi.KeteranganMasuk,
	}

}

func ToPresensiKeluarResponse(presensi domain.Presensi) web.PresensiKeluarResponse {
	return web.PresensiKeluarResponse{
		Keterangan: presensi.KeteranganKeluar,
	}

}

func ToRiwayatPresensiResponse(presensi domain.Presensi) web.RiwayatPresensiResponse {
	return web.RiwayatPresensiResponse{
		IdUser:           presensi.IdUser,
		TanggalPresensi:  presensi.TanggalPresensi,
		JamMasuk:         presensi.JamMasuk,
		JamPulang:        presensi.JamPulang,
		TanggalPulang:    string(presensi.TanggalPulang),
		KeteranganMasuk:  presensi.KeteranganMasuk,
		KeteranganKeluar: presensi.KeteranganKeluar,
		Koordinat:        presensi.Koordinat,
		Alamat:           presensi.Alamat,
		Status:           presensi.Status,
	}
}

func ToRiwayatPresensiResponses(presensi []domain.Presensi) []web.RiwayatPresensiResponse {
	var riwayatResponses []web.RiwayatPresensiResponse
	for _, presensi := range presensi {
		riwayatResponses = append(riwayatResponses, ToRiwayatPresensiResponse(presensi))
	}
	return riwayatResponses
}
