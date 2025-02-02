package mqtt

import "github.com/torbenconto/bambulabs_api/hms"

type Message struct {
	Print struct {
		Ams struct {
			Ams []struct {
				Humidity string `json:"humidity"`
				ID       string `json:"id"`
				Temp     string `json:"temp"`
				Tray     []struct {
					ID            string   `json:"id"`
					BedTemp       string   `json:"bed_temp,omitempty"`
					BedTempType   string   `json:"bed_temp_type,omitempty"`
					Cols          []string `json:"cols,omitempty"`
					DryingTemp    string   `json:"drying_temp,omitempty"`
					DryingTime    string   `json:"drying_time,omitempty"`
					NozzleTempMax string   `json:"nozzle_temp_max,omitempty"`
					NozzleTempMin string   `json:"nozzle_temp_min,omitempty"`
					Remain        int      `json:"remain,omitempty"`
					TagUID        string   `json:"tag_uid,omitempty"`
					TrayColor     string   `json:"tray_color,omitempty"`
					TrayDiameter  string   `json:"tray_diameter,omitempty"`
					TrayIDName    string   `json:"tray_id_name,omitempty"`
					TrayInfoIdx   string   `json:"tray_info_idx,omitempty"`
					TraySubBrands string   `json:"tray_sub_brands,omitempty"`
					TrayType      string   `json:"tray_type,omitempty"`
					TrayUUID      string   `json:"tray_uuid,omitempty"`
					TrayWeight    string   `json:"tray_weight,omitempty"`
					XcamInfo      string   `json:"xcam_info,omitempty"`
				} `json:"tray"`
			} `json:"ams"`
			AmsExistBits     string `json:"ams_exist_bits"`
			InsertFlag       bool   `json:"insert_flag"`
			PowerOnFlag      bool   `json:"power_on_flag"`
			TrayExistBits    string `json:"tray_exist_bits"`
			TrayIsBblBits    string `json:"tray_is_bbl_bits"`
			TrayNow          string `json:"tray_now"`
			TrayReadDoneBits string `json:"tray_read_done_bits"`
			TrayReadingBits  string `json:"tray_reading_bits"`
			TrayTar          string `json:"tray_tar"`
			Version          int    `json:"version"`
		} `json:"ams"`
		AmsRfidStatus           int            `json:"ams_rfid_status"`
		AmsStatus               int            `json:"ams_status"`
		AuxPartFan              bool           `json:"aux_part_fan"`
		BedTargetTemper         float64        `json:"bed_target_temper"`
		BedTemper               float64        `json:"bed_temper"`
		BigFan1Speed            string         `json:"big_fan1_speed"`
		BigFan2Speed            string         `json:"big_fan2_speed"`
		ChamberTemper           float64        `json:"chamber_temper"`
		Command                 string         `json:"command"`
		CoolingFanSpeed         string         `json:"cooling_fan_speed"`
		FailReason              string         `json:"fail_reason"`
		FanGear                 int            `json:"fan_gear"`
		FilamBak                []any          `json:"filam_bak"`
		ForceUpgrade            bool           `json:"force_upgrade"`
		GcodeFile               string         `json:"gcode_file"`
		GcodeFilePreparePercent string         `json:"gcode_file_prepare_percent"`
		GcodeStartTime          string         `json:"gcode_start_time"`
		GcodeState              string         `json:"gcode_state"`
		HeatbreakFanSpeed       string         `json:"heatbreak_fan_speed"`
		HMS                     []hms.HMSError `json:"hms"`
		HomeFlag                int            `json:"home_flag"`
		HwSwitchState           int            `json:"hw_switch_state"`
		Ipcam                   struct {
			IpcamDev    string `json:"ipcam_dev"`
			IpcamRecord string `json:"ipcam_record"`
			Resolution  string `json:"resolution"`
			Timelapse   string `json:"timelapse"`
		} `json:"ipcam"`
		LayerNum     int    `json:"layer_num"`
		Lifecycle    string `json:"lifecycle"`
		LightsReport []struct {
			Mode string `json:"mode"`
			Node string `json:"node"`
		} `json:"lights_report"`
		Maintain            int     `json:"maintain"`
		McPercent           int     `json:"mc_percent"`
		McPrintErrorCode    string  `json:"mc_print_error_code"`
		McPrintStage        string  `json:"mc_print_stage"`
		McPrintSubStage     int     `json:"mc_print_sub_stage"`
		McRemainingTime     int     `json:"mc_remaining_time"`
		MessProductionState string  `json:"mess_production_state"`
		NozzleDiameter      string  `json:"nozzle_diameter"`
		NozzleTargetTemper  float64 `json:"nozzle_target_temper"`
		NozzleTemper        float64 `json:"nozzle_temper"`
		Online              struct {
			Ahb     bool `json:"ahb"`
			Rfid    bool `json:"rfid"`
			Version int  `json:"version"`
		} `json:"online"`
		PrintError       int    `json:"print_error"`
		PrintGcodeAction int    `json:"print_gcode_action"`
		PrintRealAction  int    `json:"print_real_action"`
		PrintType        string `json:"print_type"`
		ProfileID        string `json:"profile_id"`
		ProjectID        string `json:"project_id"`
		QueueNumber      int    `json:"queue_number"`
		Sdcard           bool   `json:"sdcard"`
		SequenceID       string `json:"sequence_id"`
		SpdLvl           int    `json:"spd_lvl"`
		SpdMag           int    `json:"spd_mag"`
		Stg              []any  `json:"stg"`
		StgCur           int    `json:"stg_cur"`
		SubtaskID        string `json:"subtask_id"`
		SubtaskName      string `json:"subtask_name"`
		TaskID           string `json:"task_id"`
		TotalLayerNum    int    `json:"total_layer_num"`
		UpgradeState     struct {
			AhbNewVersionNumber string `json:"ahb_new_version_number"`
			AmsNewVersionNumber string `json:"ams_new_version_number"`
			ConsistencyRequest  bool   `json:"consistency_request"`
			DisState            int    `json:"dis_state"`
			ErrCode             int    `json:"err_code"`
			ForceUpgrade        bool   `json:"force_upgrade"`
			Message             string `json:"message"`
			Module              string `json:"module"`
			NewVersionState     int    `json:"new_version_state"`
			OtaNewVersionNumber string `json:"ota_new_version_number"`
			Progress            string `json:"progress"`
			SequenceID          int    `json:"sequence_id"`
			Status              string `json:"status"`
		} `json:"upgrade_state"`
		Upload struct {
			FileSize      int    `json:"file_size"`
			FinishSize    int    `json:"finish_size"`
			Message       string `json:"message"`
			OssURL        string `json:"oss_url"`
			Progress      int    `json:"progress"`
			SequenceID    string `json:"sequence_id"`
			Speed         int    `json:"speed"`
			Status        string `json:"status"`
			TaskID        string `json:"task_id"`
			TimeRemaining int    `json:"time_remaining"`
			TroubleID     string `json:"trouble_id"`
		} `json:"upload"`
		VtTray struct {
			BedTemp       string   `json:"bed_temp"`
			BedTempType   string   `json:"bed_temp_type"`
			Cols          []string `json:"cols"`
			DryingTemp    string   `json:"drying_temp"`
			DryingTime    string   `json:"drying_time"`
			ID            string   `json:"id"`
			NozzleTempMax string   `json:"nozzle_temp_max"`
			NozzleTempMin string   `json:"nozzle_temp_min"`
			Remain        int      `json:"remain"`
			TagUID        string   `json:"tag_uid"`
			TrayColor     string   `json:"tray_color"`
			TrayDiameter  string   `json:"tray_diameter"`
			TrayIDName    string   `json:"tray_id_name"`
			TrayInfoIdx   string   `json:"tray_info_idx"`
			TraySubBrands string   `json:"tray_sub_brands"`
			TrayType      string   `json:"tray_type"`
			TrayUUID      string   `json:"tray_uuid"`
			TrayWeight    string   `json:"tray_weight"`
			XcamInfo      string   `json:"xcam_info"`
		} `json:"vt_tray"`
		WifiSignal string `json:"wifi_signal"`
		Xcam       struct {
			AllowSkipParts           bool   `json:"allow_skip_parts"`
			BuildplateMarkerDetector bool   `json:"buildplate_marker_detector"`
			FirstLayerInspector      bool   `json:"first_layer_inspector"`
			HaltPrintSensitivity     string `json:"halt_print_sensitivity"`
			PrintHalt                bool   `json:"print_halt"`
			PrintingMonitor          bool   `json:"printing_monitor"`
			SpaghettiDetector        bool   `json:"spaghetti_detector"`
		} `json:"xcam"`
		XcamStatus string `json:"xcam_status"`
	} `json:"print"`
}
