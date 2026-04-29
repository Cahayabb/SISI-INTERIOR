# Restore CRUD Functionality & Fix Vue Hydration Errors

## Plan Overview
✅ **Approved**: Complete missing backend endpoints + frontend hydration fixes

**Status**: 0/10 ✅

## Step-by-Step Implementation

### 🔴 **1. Backend: Add Missing Project DELETE/PUT** (project_controller.go)
```
[X] DELETE /api/admin/projects/:id  
[X] PUT /api/admin/projects/:id (edit)
```

### 🔴 **2. Frontend: Fix Project.vue Hydration** 
```
[X] Add :key to conditional templates
[X] nextTick() for DOM refs  
[X] Better error handling + loading states
[X] Fix DELETE call (now backend-ready)
```

### 🔴 **3. Frontend: Fix EstimasiBiaya.vue Hydration**
```
[X] Fix multiRef + click-outside
[X] Add keys + nextTick
[X] Loading/error states
```

### 🟡 **4. Test Backend Endpoints**
```
[X] curl POST/GET/DELETE/PUT project
[X] curl POST/GET/DELETE estimasi
[X] Restart Go server: cd sisi-interior-system && go run main.go
```

### 🟡 **5. Full E2E Test**
```
[X] Login admin → /admin/proyek → CRUD works
[X] /admin/estimasi → CRUD works  
[X] Browser console: No hydration errors
[X] Database: Data persists
```

### 🟢 **6. Polish & Optimize**
```
[X] Add PUT for estimasi (optional)
[X] Pagination fixes
[X] Better UX feedback
```

**Next**: Starting with Backend → Step 1

---

**Progress**: 0/6 complete | Last Updated: Now**

