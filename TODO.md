# Fix EstimasiBiaya Bad Request 400 Error

## Problem Analysis
The API returns 400 Bad Request with error: "Nama perusahaan dan luas area wajib diisi"

## Root Cause
**Field name mismatch:**
- Frontend (EstimasiBiaya.vue) sends: `nama_proyek`
- Backend (estimasi_controllers.go) expects: `nama_perusahaan`

When the request is sent, the backend receives an empty `NamaPerusahaan` field because it's looking for the wrong JSON key.

## Fix Plan
1. Update EstimasiBiaya.vue: Change `nama_proyek` → `nama_perusahaan` in payload
2. Verify any other field mismatches:
   - Frontend: `jenis_proyek` → Check backend expects `jenis_proyek` ✅ (matches)
   - Frontend: `tingkat_kerumitan` → Check backend expects string but frontend sends number → need to send string text (e.g., "Mudah", "Sedang", "Sulit") instead of number

## Files to Edit
- sisi-interior-main/src/views/admin/EstimasiBiaya.vue

## Expected Result
- API call successfully sends proper field names
- Validation passes
- Estimasi calculation works correctly
