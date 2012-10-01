package kdfutil

import (
	"crypto/sha256" 
	"testing" 
	"time" 
) 
 
func TestCalibrate(t *testing.T) { 
        if iter := CalibratePBKDF2(10*time.Minute, sha256.New, 64); iter < 1 { 
                t.Errorf("calibration failed: got %d iterations for 10 minutes, expecting at least 1", iter) 
        } 
} 

